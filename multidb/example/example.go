package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb"
)

const (
	userDSN = "root:password@tcp(127.0.0.1:3306)/userdb"
	itemDSN = "root:password@tcp(127.0.0.1:3306)/itemdb"
)

func Run() {
	if err := run(); err != nil {
		log.Printf("Error while run: %v", err)
		os.Exit(1)
	}
}

func run() (retErr error) {
	ctx := context.Background()

	db, err := multidb.Connect(ctx, userDSN, itemDSN)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	// use userDB
	userId := 1
	user, err := db.GetUser(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	fmt.Printf("Got user: %+v\n", user)

	// use itemDB
	itemId := 2
	item, err := db.GetItem(ctx, itemId)
	if err != nil {
		return fmt.Errorf("failed to get item: %w", err)
	}

	fmt.Printf("Got item: %+v\n", item)

	// use transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if retErr != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("Failed to rollback transaction: %v", err)
			}
		}
	}()

	userName := "Usagi"
	newUser, err := tx.AddUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	fmt.Printf("Added user: %+v\n", newUser)

	itemName := "Sukiyaki"
	newItem, err := tx.AddItem(ctx, itemName)
	if err != nil {
		return fmt.Errorf("failed to add item: %w", err)
	}

	fmt.Printf("Added item: %+v\n", newItem)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
