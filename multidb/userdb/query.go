package userdb

import (
	"context"

	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/common"
)

func getUser(ctx context.Context, getter common.GetterCtx, id int) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = ?"
	if err := getter.GetContext(ctx, &user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func addUser(ctx context.Context, execer common.ExecerCtx, name string) (*User, error) {
	query := "INSERT INTO users (name) VALUES (?)"
	result, err := execer.ExecContext(ctx, query, name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return NewUser(int(id), name), nil
}
