package db

type PgTaskStore struct {
	store *PostgresDB
}

func NewPgTaskStore(db *PostgresDB) (*PgTaskStore, error) {
	return &PgTaskStore{store: db}, nil
}

