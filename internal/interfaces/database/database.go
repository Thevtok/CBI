package database

import (
	"database/sql"
)

type DB interface {
	Begin() (*sql.Tx, error)
	Commit() error
	Rollback() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

type Database struct {
	Conn *sql.DB
	Tx   *sql.Tx
}

func (d *Database) Begin() (*sql.Tx, error) {
	return d.Conn.Begin()
}

func (d *Database) Commit() error {
	return d.Tx.Commit()
}

func (d *Database) Rollback() error {
	return d.Tx.Rollback()
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Conn.Exec(query, args...)
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.Conn.Query(query, args...)
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.Conn.QueryRow(query, args...)
}

func (d *Database) Prepare(query string) (*sql.Stmt, error) {
	return d.Conn.Prepare(query)
}
