package multidb

import (
	"context"
	"fmt"
	"log"

	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/itemdb"
	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/userdb"
)

type DB struct {
	*userdb.UserDB
	*itemdb.ItemDB
}

func NewDB(userDB *userdb.UserDB, itemDB *itemdb.ItemDB) *DB {
	return &DB{
		UserDB: userDB,
		ItemDB: itemDB,
	}
}

const driver = "mysql"

func Connect(ctx context.Context, userDSN, itemDSN string) (*DB, error) {
	userDB, err := userdb.Connect(ctx, driver, userDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect userDB: %w", err)
	}

	itemDB, err := itemdb.Connect(ctx, driver, itemDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect itemDB: %w", err)
	}

	return NewDB(userDB, itemDB), nil
}

func (db *DB) Close() {
	if err := db.UserDB.Close(); err != nil {
		log.Printf("Failed to close userDB: %v", err)
	}
	if err := db.ItemDB.Close(); err != nil {
		log.Printf("Failed to close itemDB: %v", err)
	}
}

func (db *DB) Begin() (*Tx, error) {
	userTx, err := db.UserDB.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin userDB: %w", err)
	}

	itemTx, err := db.ItemDB.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin itemDB: %w", err)
	}

	return newTx(userTx, itemTx), nil
}
