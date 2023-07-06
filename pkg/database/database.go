package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnection struct {
	Conn *pgxpool.Pool
}

func (db *DBConnection) InitConnection(db_conn_str string) error {
	var err error
	db.Conn, err = pgxpool.New(context.Background(), db_conn_str)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnection) Close() {
	db.Conn.Close()
}
