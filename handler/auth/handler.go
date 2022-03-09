package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Check ...
// @Summary check api
// @Description check integrating swagger
// @Tags Authentication
// @Accept json
// @Success 200
// @Router /api/v1/auth/temp [get]
func Check(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "data": "ok",
    })
}
