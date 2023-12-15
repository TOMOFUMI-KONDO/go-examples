package itemdb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ItemTx struct {
	tx *sqlx.Tx
}

func newItemTx(tx *sqlx.Tx) *ItemTx {
	return &ItemTx{
		tx: tx,
	}
}

func (t *ItemTx) Commit() error {
	return t.tx.Commit()
}

func (t *ItemTx) Rollback() error {
	return t.tx.Rollback()
}

func (t *ItemTx) GetItem(ctx context.Context, id int) (*Item, error) {
	return getItem(ctx, t.tx, id)
}

func (t *ItemTx) AddItem(ctx context.Context, name string) (*Item, error) {
	return addItem(ctx, t.tx, name)
}
