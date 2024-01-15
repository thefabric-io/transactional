package pgtransactional

import (
	"context"
	"database/sql"
	"github.com/thefabric-io/transactional"
	"runtime"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSQLXTransactionalConnection(ctx context.Context, connectionString string) (*SQLXTransactional, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLXTransactional{db: db}, nil
}

type SQLXTransactional struct {
	db *sqlx.DB
}

func (t *SQLXTransactional) Ping(ctx context.Context) error {
	return t.db.PingContext(ctx)
}

func (t *SQLXTransactional) DB() *sqlx.DB {
	return t.db
}

func (t *SQLXTransactional) DefaultLogFields() map[string]any {
	_, file, _, _ := runtime.Caller(0)
	return map[string]any{
		"metadata": map[string]any{
			"file": file,
		},
		"subject": map[string]any{
			"type":       "SQLXTransactional",
			"implements": "Transactional",
		},
	}
}

func (t *SQLXTransactional) BeginTransaction(ctx context.Context, opts transactional.BeginTransactionOptions) (transactional.Transaction, error) {
	isolationLevel := IsolationLevel(opts.IsolationLevel)
	isReadOnly := IsReadOnly(opts.AccessMode)

	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: isolationLevel,
		ReadOnly:  isReadOnly,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func IsolationLevel(level transactional.TxIsoLevel) sql.IsolationLevel {
	switch level {
	case transactional.Serializable:
		return sql.LevelSerializable
	case transactional.RepeatableRead:
		return sql.LevelRepeatableRead
	case transactional.ReadCommitted:
		return sql.LevelReadCommitted
	case transactional.ReadUncommitted:
		return sql.LevelReadUncommitted
	}

	return sql.LevelDefault
}

func IsReadOnly(level transactional.TxAccessMode) bool {
	if level == transactional.ReadOnly {
		return true
	}

	return false
}
