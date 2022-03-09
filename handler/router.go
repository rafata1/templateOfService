package handler

import (
    "github.com/gin-gonic/gin"
    authHandler "github.com/templateOfService/handler/auth"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    apiV1 := router.Group("/api/v1")
    {
        auth := apiV1.Group("/auth")
        {
            auth.GET("/temp", authHandler.Check)
        }
    }
    return router
}
