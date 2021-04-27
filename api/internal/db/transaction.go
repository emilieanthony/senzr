package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	tx *sqlx.Tx
	db *sqlx.DB
}

func (t *Transaction) Get(dest interface{}, query string) error {
	err := t.tx.Get(dest, query)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Select(dest interface{}, query string) error {
	err := t.tx.Select(dest, query)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return t.tx.NamedExec(query, arg)
}

func (t *Transaction) Query(query string) (*sqlx.Rows, error) {
	rows, err := t.tx.Queryx(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (t *Transaction) Commit() error {
	if t.db == nil {
		return fmt.Errorf("connection is already closed")
	}
	err := t.tx.Commit()
	if err != nil {
		return err
	}
	err = t.db.Close()
	if err != nil {
		return err
	}
	t.tx = nil
	t.db = nil
	return nil
}

func BeginTransaction() (_ *Transaction, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("transaction: %v", err)
		}
	}()
	db, err := newDBConnection()
	if err != nil {
		return nil, err
	}
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return &Transaction{
		tx: tx,
		db: db,
	}, nil
}
