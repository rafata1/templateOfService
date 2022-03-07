package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// @BasePath /api/v1/auth

// Signup godoc
// @Description signup api
// @Accept json
// @Produce json
// @Router /signup [get]
func SignUp(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "signed up",
    })
    return
}
