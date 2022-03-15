package mysql

import (
    "github.com/jmoiron/sqlx"
    "os"
)

type MySQL struct {
    conn *sqlx.DB
    url  string
}

var mysql *MySQL

func Connect() error {
    url := os.Getenv("MYSQL")
    conn, err := sqlx.Connect("mysql", url)
    mysql = &MySQL{
        conn: conn,
        url:  url,
    }
    return err
}

func GetMySQLInstance() *sqlx.DB {
    return mysql.conn
}
