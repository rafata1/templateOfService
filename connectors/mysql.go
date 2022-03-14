package connectors

import (
    "github.com/jmoiron/sqlx"
    "log"
    "os"
)

var conn *sqlx.DB

func ConnectMySQL() *sqlx.DB {
    if conn == nil {
        var err error
        conn, err = sqlx.Connect("mysql", os.Getenv("MYSQL"))
        if err != nil {
            log.Printf("Error connecting to mysql: %s", err)
        }
    }
    return conn
}
