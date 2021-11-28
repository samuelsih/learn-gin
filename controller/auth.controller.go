package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController is a interface to control the authentication / login
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

//authController is private struct to store/put the service
type authController struct {

}

//NewAuthController creates new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "login section",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "register section",
	})
} 