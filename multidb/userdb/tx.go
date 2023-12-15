package userdb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserTx struct {
	tx *sqlx.Tx
}

func newUserTx(tx *sqlx.Tx) *UserTx {
	return &UserTx{
		tx: tx,
	}
}

func (t *UserTx) Commit() error {
	return t.tx.Commit()
}

func (t *UserTx) Rollback() error {
	return t.tx.Rollback()
}

func (t *UserTx) GetUser(ctx context.Context, id int) (*User, error) {
	return getUser(ctx, t.tx, id)
}

func (t *UserTx) AddUser(ctx context.Context, name string) (*User, error) {
	return addUser(ctx, t.tx, name)
}
