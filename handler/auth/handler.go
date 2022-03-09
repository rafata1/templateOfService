package auth

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Signup ...
//// @Tags         Authentication
//// @Accept       json
//// @Produce      json
//// @Param          {object} SignupReq
//// @Success      200  {object}  BaseRes
//// @Router       /api/v1/auth/signup [post]
func Signup(c *gin.Context) {
    var req SignupReq
    err := c.BindJSON(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "message format is invalid",
        })
        return
    }
    fmt.Println(req)
    c.JSON(http.StatusOK, gin.H{
        "message": "received",
    })
}
