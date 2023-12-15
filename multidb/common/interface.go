package common

import (
	"context"
	"database/sql"
)

type GetterCtx interface {
	GetContext(ctx context.Context, dest any, query string, args ...any) error
}

type ExecerCtx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
