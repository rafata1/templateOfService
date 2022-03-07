package main

import (
    "github.com/joho/godotenv"
    "github.com/templateOfService/handler"
    "log"
    "sync"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env")
    }
    h := handler.NewHandler()
    var wg sync.WaitGroup
    wg.Add(1)
    h.Run(&wg)
    wg.Wait()
}
