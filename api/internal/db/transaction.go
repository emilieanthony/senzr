package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Db struct{}

type Database interface {
	BeginTransaction() (_ Transaction, err error)
}

var _ Database = &Db{}

type Transaction interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Query(query string) (*sqlx.Rows, error)
	Commit() error
}

type Tx struct {
	Tx *sqlx.Tx
	db *sqlx.DB
}

var _ Transaction = &Tx{}

func (t *Tx) Get(dest interface{}, query string, args ...interface{}) error {
	var err error
	if len(args) != 0 {
		err = t.Tx.Get(&dest, query, args)
	} else {
		err = t.Tx.Get(&dest, query)
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *Tx) Select(dest interface{}, query string, args ...interface{}) error {
	var err error
	if len(args) != 0 {
		err = t.Tx.Select(dest, query, args)
	} else {
		err = t.Tx.Select(dest, query)
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *Tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return t.Tx.NamedExec(query, arg)
}

func (t *Tx) Query(query string) (*sqlx.Rows, error) {
	rows, err := t.Tx.Queryx(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (t *Tx) Commit() error {
	if t.db == nil {
		return fmt.Errorf("connection is already closed")
	}
	err := t.Tx.Commit()
	if err != nil {
		return err
	}
	err = t.db.Close()
	if err != nil {
		return err
	}
	t.Tx = nil
	t.db = nil
	return nil
}

func (d *Db) BeginTransaction() (_ Transaction, err error) {
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
	return &Tx{
		Tx: tx,
		db: db,
	}, nil
}
