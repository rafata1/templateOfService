package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/templateOfService/model"
    "github.com/templateOfService/service/auth"
    "net/http"
    "strings"
)

type Handler struct {
    service *auth.Service
}

func NewHandler() *Handler {
    return &Handler{
        service: auth.NewService(),
    }
}

// Signup godoc
// @Summary      Register an account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        signupRequest  body SignupReq true "Signup"
// @Success      200 {object} BaseRes
// @Router       /api/v1/auth/signup [post]
func (h *Handler) Signup(c *gin.Context) {
    var req SignupReq
    err := c.BindJSON(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, BaseRes{
            Message: "request format is not correct",
        })
        return
    }
    user := transformSignupRequestToUserModel(req)
    message := h.service.Signup(user)
    res := BaseRes{
        Message: message,
    }
    c.JSON(http.StatusOK, res)
}

func transformSignupRequestToUserModel(req SignupReq) model.User {
    return model.User{
        Username: strings.TrimSpace(req.Username),
        Email:    strings.TrimSpace(req.Email),
        Password: strings.TrimSpace(req.Password),
        Phone:    strings.TrimSpace(req.Phone),
    }
}
