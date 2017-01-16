package transactor

import (
	"database/sql"
	. "github.com/MakeNowJust/heredoc/dot"
	_ "github.com/lib/pq"
	"strings"
	"fmt"
)

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

func (transactor *Connection) buildInsert(entityID uint64, attributeValues AttributeValues) error {
	insertStatement := `INSERT INTO eavt (entity, attribute, value, time) VALUES`

	params := []interface{}{}

	for attribute, value := range attributeValues {
		insertStatement += fmt.Sprintf(
			"(%d, $%d, $%d, current_timestamp),\n",
			entityID,
			attribute,
			value,
		)
		params = append(params, attribute, value)

	}

	insertStatement = strings.TrimSuffix(insertStatement, ",\n")
	insertStatement += ";\n"

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

func (transactor *Connection) UpdateEntity(entityID uint64, attributeValues AttributeValues) error {
	if err := transactor.buildInsert(entityID, attributeValues); err != nil {
		return err
	}

	return nil
}

type AttributeValues map[string]interface{}

type Entity struct {
	Id uint64
	AttributeValues AttributeValues
}

func (transactor *Connection) CreateEntity(attributeValues AttributeValues) (uint64, error) {
	var entityId uint64

	if err := transactor.DB.QueryRow("select nextval('entity_sequence');").Scan(&entityId); err != nil {
		return nil, err
	}

	if err := transactor.buildInsert(entityId, attributeValues); err != nil {
		return nil, err
	}

	return entityId, nil
}

func ensureSchema(db *sql.DB) error {
	_, err := db.Exec(
		D(`
		  CREATE TABLE IF NOT EXISTS eavt (
		    entity bigint NOT NULL,
		    attribute text NOT NULL,
		  	value text,
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
