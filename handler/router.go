package handler

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "github.com/templateOfService/handler/auth"
    "log"
)

func SetupRouter() {
    authHandler := auth.NewHandler()
    router := gin.Default()
    apiV1 := router.Group("/api/v1")
    {
        authRouter := apiV1.Group("/auth")
        {
            authRouter.POST("/signup", authHandler.Signup)
        }
    }
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    err := router.Run()
    if err != nil {
        log.Fatalf("Error starting server: %s", err.Error())
    }
}
