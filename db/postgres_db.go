package db

import (
	"database/sql"
	"os"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(dbString string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dbString)

	if err != nil {
		return nil, err
	}

	database := &PostgresDB{db: db}

	if err := database.createSchema(); err != nil {
		return nil, err
	}

	return database, nil
}

func (s *PostgresDB) createSchema() error {
	schema, err := os.ReadFile("sql/create_schema.sql")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(string(schema))
	return err
}

func (s *PostgresDB) close(){
	s.close()
}