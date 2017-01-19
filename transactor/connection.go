package transactor

import (
	"database/sql"
	. "github.com/MakeNowJust/heredoc/dot"
	_ "github.com/lib/pq"
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
