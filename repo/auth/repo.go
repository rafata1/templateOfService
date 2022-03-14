package authRepo

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/templateOfService/model"
    "github.com/templateOfService/repo"
)

type Repo struct {
    db *sqlx.DB
}

func NewRepo() *Repo {
    return &Repo{
        db: repo.GetMysqlInstance(),
    }
}

func (r *Repo) InsertUser(user model.User) error {
    fmt.Println(user)
    return nil
}
