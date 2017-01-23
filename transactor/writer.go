package transactor

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

// string
type AttributeValues map[string]json.RawMessage

func (transactor *Connection) insert(entityID uint64, attributeValues AttributeValues) error {
	insertStatement := `INSERT INTO eavt (entity, attribute, json_value, time) VALUES`

	params := []interface{}{}

	var parameterIndex = 1

	for attribute, value := range attributeValues {
		insertStatement += fmt.Sprintf(
			"(%d, $%d, $%d, current_timestamp),\n",
			entityID,
			parameterIndex,
			parameterIndex+1,
		)

		params = append(params, attribute, string(value))

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

func (transactor *Connection) UpdateEntity(entityID uint64, attributeValues AttributeValues) error {
	if err := transactor.insert(entityID, attributeValues); err != nil {
		return err
	}

	return nil
}

func (transactor *Connection) CreateEntity(attributeValues AttributeValues) (uint64, error) {
	var entityId uint64

	if err := transactor.DB.QueryRow("select nextval('entity_sequence');").Scan(&entityId); err != nil {
		return 0, err
	}

	if err := transactor.insert(entityId, attributeValues); err != nil {
		return 0, err
	}

	return entityId, nil
}
