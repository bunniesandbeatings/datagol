package transactor

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"github.com/bunniesandbeatings/datagol/attributes"
	"strconv"
	"time"
)

type StringEntity map[string]string

func (backend *Backend) insert(entityId uint64, entity StringEntity) error {
	insertStatement := `INSERT INTO eavt (entity, attribute, value, time) VALUES`
	params := []interface{}{}

	var parameterIndex = 1

	for attribute, value := range entity {
		insertStatement += fmt.Sprintf(
			"(%d, $%d, $%d, current_timestamp),\n",
			entityId,
			parameterIndex,
			parameterIndex+1,
		)

		params = append(params, attribute, string(value))

		parameterIndex = parameterIndex + 2
	}

	insertStatement = strings.TrimSuffix(insertStatement, ",\n")
	insertStatement += " RETURNING current_timestamp;\n"

	//log.Printf("DEBUG: Insert Statement: %s", insertStatement)
	//log.Printf("DEBUG: Insert params: %v", params)

	preparedStatement, err := backend.DB.Prepare(insertStatement)
	if err != nil {
		return err
	}

	var assertedAt time.Time;

	err = preparedStatement.QueryRow(params...).Scan(&assertedAt)

	if err != nil {
		return err
	}

	preparedStatement.Close()

	entity[attributes.IDENT] = strconv.FormatUint(entityId,10)
	entity[attributes.TIME] = assertedAt.Format(time.RFC3339Nano)

	return nil
}

func (backend *Backend) Commit(entity StringEntity) (StringEntity, error) {
	var entityId uint64

	if err := backend.DB.QueryRow("select nextval('entity_sequence');").Scan(&entityId); err != nil {
		return entity, err
	}

	if err := backend.insert(entityId, entity); err != nil {
		return entity, err
	}

	entity[attributes.IDENT] = strconv.FormatUint(entityId,10)
	return entity, nil
}
