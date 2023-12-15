package userdb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func Connect(ctx context.Context, dsn, driver string) (*UserDB, error) {
	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, err
	}

	return NewUserDB(db), nil
}

func (d *UserDB) Close() error {
	return d.db.Close()
}

func (d *UserDB) Begin() (*UserTx, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	return newUserTx(tx), nil
}

func (d *UserDB) GetUser(ctx context.Context, id int) (*User, error) {
	return getUser(ctx, d.db, id)
}

func (d *UserDB) AddUser(ctx context.Context, name string) (*User, error) {
	return addUser(ctx, d.db, name)
}
