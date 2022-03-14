package main

import (
    "github.com/joho/godotenv"
    _ "github.com/templateOfService/docs"
    "github.com/templateOfService/route"
    "log"
)

// @title User API documentation
// @BasePath /
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading env: %s", err.Error())
    }

    router := route.InitRouter()
    err = router.Run()
    if err != nil {
        log.Fatalf("Error starting server: %s", err.Error())
    }
}
