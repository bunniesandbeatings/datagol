package transactor

import (
	"database/sql"
	. "github.com/MakeNowJust/heredoc/dot"
	_ "github.com/lib/pq"
	"strings"
	"fmt"
	"log"
	"encoding/json"
)

type AttributeValuesJson map[string]json.RawMessage

type Entity struct {
	Id              uint64
	AttributeValues AttributeValuesJson
}

type Connection struct {
	DB *sql.DB
}

func NewConnection(dataSourceName string) (*Connection, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := ensureSchema(db); err != nil {
		return nil, err
	}

	return &Connection{DB: db}, nil
}

func (transactor *Connection) insert(entityID uint64, attributeValues AttributeValuesJson) error {
	insertStatement := `INSERT INTO eavt (entity, attribute, json_value, time) VALUES`

	params := []interface{}{}

	var parameterIndex = 1

	for attribute, value := range attributeValues {
		insertStatement += fmt.Sprintf(
			"(%d, $%d, $%d, current_timestamp),\n",
			entityID,
			parameterIndex,
			parameterIndex + 1,
		)

		params = append(params, attribute, string(value)	)

		parameterIndex = parameterIndex + 2
	}

	insertStatement = strings.TrimSuffix(insertStatement, ",\n")
	insertStatement += ";\n"

	log.Printf("DEBUG: Insert Statement: %s", insertStatement)
	log.Printf("DEBUG: Insert params: %v", params)

	preparedStatement, err := transactor.DB.Prepare(insertStatement)
	if err != nil {
		return err
	}

	_, err = preparedStatement.Exec(params...)
	if err != nil {
		return err
	}

	preparedStatement.Close()

	return nil
}

func (transactor *Connection) UpdateEntity(entityID uint64, attributeValues AttributeValuesJson) error {
	if err := transactor.insert(entityID, attributeValues); err != nil {
		return err
	}

	return nil
}

func (transactor *Connection) CreateEntity(attributeValues AttributeValuesJson) (uint64, error) {
	var entityId uint64

	if err := transactor.DB.QueryRow("select nextval('entity_sequence');").Scan(&entityId); err != nil {
		return 0, err
	}

	if err := transactor.insert(entityId, attributeValues); err != nil {
		return 0, err
	}

	return entityId, nil
}

func ensureSchema(db *sql.DB) error {
	_, err := db.Exec(
		D(`
		  CREATE TABLE IF NOT EXISTS eavt (
		    entity bigint NOT NULL,
		    attribute text NOT NULL,
		    json_value text NOT NULL,
		  	time timestamp NOT NULL
		  );
		  CREATE SEQUENCE IF NOT EXISTS entity_sequence;
		  CREATE INDEX IF NOT EXISTS enitity_index ON eavt(entity);
		  CREATE INDEX IF NOT EXISTS time_index ON eavt(time DESC);
		`),
	)

	if err != nil {
		return err
	}

	return nil
}
