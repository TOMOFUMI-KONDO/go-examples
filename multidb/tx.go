package multidb

import (
	"errors"
	"log"

	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/itemdb"
	"github.com/TOMOFUMI-KONDO/go-sandbox/multidb/userdb"
)

type Tx struct {
	*userdb.UserTx
	*itemdb.ItemTx
}

func newTx(userTx *userdb.UserTx, itemTx *itemdb.ItemTx) *Tx {
	return &Tx{
		UserTx: userTx,
		ItemTx: itemTx,
	}
}

func (tx *Tx) Commit() error {
	var errs []error

	if err := tx.UserTx.Commit(); err != nil {
		errs = append(errs, err)
	}
	if err := tx.ItemTx.Commit(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (tx *Tx) Rollback() error {
	var errs []error

	if err := tx.UserTx.Rollback(); err != nil {
		errs = append(errs, err)
	}
	if err := tx.ItemTx.Rollback(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (tc *Tx) CommitOrRollback(err error) {
	if err != nil {
		if errRB := tc.Rollback(); errRB != nil {
			log.Printf("Failed to rollback: %v", errRB)
		}
		return
	}

	if errCmt := tc.Commit(); errCmt != nil {
		if errRB := tc.Rollback(); errRB != nil {
			log.Printf("Failed to rollback: %v", errRB)
		}
		log.Printf("Failed to commit: %v", errCmt)
	}
}
