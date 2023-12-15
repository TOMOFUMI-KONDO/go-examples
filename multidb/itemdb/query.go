package itemdb

import (
	"context"
	"fmt"

	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/common"
)

func getItem(ctx context.Context, getter common.GetterCtx, id int) (*Item, error) {
	var item Item
	query := "SELECT * FROM items WHERE id = ?"
	if err := getter.GetContext(ctx, item, query, id); err != nil {
		return nil, err
	}

	return &item, nil
}

func addItem(ctx context.Context, execer common.ExecerCtx, name string) (*Item, error) {
	query := "INSERT TO items (name) VALUES (?)"
	result, err := execer.ExecContext(ctx, query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get ID of created item: %w", err)
	}

	return NewItem(int(id), name), nil
}
