package handler

import (
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/swag/example/basic/docs"
    authHandler "github.com/templateOfService/handler/auth"
    "log"
    "sync"
)

type Handler struct {
    router *gin.Engine
}

func (h Handler) Run(wg *sync.WaitGroup) {
    defer wg.Done()
    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := h.router.Group("/api/v1")
    {
        auth := v1.Group("/auth")
        {
            auth.GET("/signup", authHandler.SignUp)
        }
    }
    h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    err := h.router.Run()
    if err != nil {
        log.Fatalf("start server error: %s", err.Error())
        return
    }
    log.Println("listening on 0.0.0.0:8000")
}

func NewHandler() *Handler {
    return &Handler{
        router: gin.Default(),
    }
}
