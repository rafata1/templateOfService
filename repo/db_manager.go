package repo

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "log"
    "os"
)

type DBManager struct {
    mysqlConn *sqlx.DB
}

var manager DBManager

func SetupDB() {
    mysqlSrc := os.Getenv("MYSQL")
    fmt.Println(mysqlSrc)
    mysqlConn, err := sqlx.Connect("mysql", mysqlSrc)
    if err != nil {
        log.Fatalf("Error connecting to mysql %s", err.Error())
    }
    manager = DBManager{
        mysqlConn: mysqlConn,
    }
}

func GetMysqlInstance() *sqlx.DB {
    return manager.mysqlConn
}
