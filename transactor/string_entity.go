package transactor

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

type StringEntity map[string]string

func (backend *Backend) insert(entityID uint64, entity *StringEntity) error {
	insertStatement := `INSERT INTO eavt (entity, attribute, value, time) VALUES`

	params := []interface{}{}

	var parameterIndex = 1

	for attribute, value := range *entity {
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

	//log.Printf("DEBUG: Insert Statement: %s", insertStatement)
	//log.Printf("DEBUG: Insert params: %v", params)

	preparedStatement, err := backend.DB.Prepare(insertStatement)
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

func (backend *Backend) Commit(entity *StringEntity) (uint64, error) {
	var entityId uint64

	if err := backend.DB.QueryRow("select nextval('entity_sequence');").Scan(&entityId); err != nil {
		return 0, err
	}

	if err := backend.insert(entityId, entity); err != nil {
		return 0, err
	}

	return entityId, nil
}
