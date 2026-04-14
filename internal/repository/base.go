package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

const defaultTimeout = 3 * time.Second

type baseRepository struct {
	db *sqlx.DB
}

func newBaseRepo(db *sqlx.DB) *baseRepository {
	return &baseRepository{db: db}
}

func (r *baseRepository) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, defaultTimeout)
}

func (r *baseRepository) exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error) {
	return r.db.ExecContext(ctx, q, args...)
}

func (r *baseRepository) get(ctx context.Context, dest any, q string, args ...interface{}) error {
	return r.db.GetContext(ctx, dest, q, args...)
}

func (r *baseRepository) selectRows(ctx context.Context, dest any, q string, args ...interface{}) error {
	return r.db.SelectContext(ctx, dest, q, args...)
}

func (r *baseRepository) namedExec(ctx context.Context, q string, arg any) (sql.Result, error) {
	return r.db.NamedExecContext(ctx, q, arg)
}
