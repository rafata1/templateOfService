package route

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "github.com/templateOfService/services/auth"
)

func InitRouter() *gin.Engine {
    router := gin.Default()
    authHandler := auth.NewHandler()
    router.POST("/api/v1/auth/signup", authHandler.Signup)
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    return router
}
