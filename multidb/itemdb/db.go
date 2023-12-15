package itemdb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ItemDB struct {
	db *sqlx.DB
}

func NewItemDB(db *sqlx.DB) *ItemDB {
	return &ItemDB{
		db: db,
	}
}

func Connect(ctx context.Context, dsn, driver string) (*ItemDB, error) {
	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, err
	}

	return NewItemDB(db), nil
}

func (d *ItemDB) Close() error {
	return d.db.Close()
}

func (d *ItemDB) Begin() (*ItemTx, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	return newItemTx(tx), nil
}

func (d *ItemDB) GetItem(ctx context.Context, id int) (*Item, error) {
	return getItem(ctx, d.db, id)
}

func (d *ItemDB) AddItem(ctx context.Context, name string) (*Item, error) {
	return addItem(ctx, d.db, name)
}
