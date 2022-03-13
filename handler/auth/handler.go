package authHandler

import (
	"github.com/gin-gonic/gin"
	authCore "github.com/templateOfService/core/auth"
	"github.com/templateOfService/model"
	"net/http"
	"strings"
)

// Signup godoc
// @Summary      Register an account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        signupRequest  body SignupReq true "Signup"
// @Success      200 {object} BaseRes
// @Router       /api/v1/auth/signup [post]
func Signup(c *gin.Context) {
	var req SignupReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: "failed",
		})
		return
	}
	user := model.User{
		Username: strings.TrimSpace(req.Username),
		Email:    strings.TrimSpace(req.Email),
		Password: strings.TrimSpace(req.Password),
		Phone:    strings.TrimSpace(req.Phone),
	}
	message := authCore.Signup(user)
	res := BaseRes{
		Message: message,
	}
	c.JSON(http.StatusOK, res)
}
