package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelsih/learn-gofiber/config"
	"github.com/samuelsih/learn-gofiber/controller"
	"gorm.io/gorm"
)

//global variable
var (
	db 			   *gorm.DB  				 = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth") 
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	
	server.Run()
}