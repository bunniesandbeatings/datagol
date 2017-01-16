package util

import (
	"database/sql"
	. "github.com/MakeNowJust/heredoc/dot"
)

type Schema struct {
	DB *sql.DB
}

func (schema *Schema) TableExists(tableName string) (bool, error) {
	var hasTable bool

	err := schema.DB.QueryRow(
		D(`
				SELECT EXISTS (
					SELECT 1
					FROM information_schema.tables
					WHERE table_name = $1
				);
			`),
		tableName,
	).Scan(&hasTable)

	if err != nil {
		return false, err
	}

	return hasTable, nil
}
