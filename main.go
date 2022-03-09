package main

import (
    "github.com/joho/godotenv"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "github.com/templateOfService/docs"
    "github.com/templateOfService/handler"
    "log"
)

// @title User API documentation
// @BasePath /
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env")
    }

    router := handler.SetupRouter()
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    err = router.Run()
    if err != nil {
        log.Fatalf("Error starting server: %s", err.Error())
    }
}
