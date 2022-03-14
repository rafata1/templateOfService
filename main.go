package main

import (
    "github.com/joho/godotenv"
    _ "github.com/templateOfService/docs"
    "github.com/templateOfService/handler"
    "github.com/templateOfService/repo"
    "log"
)

// @title User API documentation
// @BasePath /
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env")
    }
    handler.SetupRouter()
    repo.SetupDB()
}
