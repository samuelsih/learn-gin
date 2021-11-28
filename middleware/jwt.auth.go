package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/samuelsih/learn-gofiber/helper"
	"github.com/samuelsih/learn-gofiber/service"
)

//AuthorizeJWT validates the token user give, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {	
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.MakeErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader) 

		if err != nil {
			panic(err.Error())
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] : ", claims["user_id"])
			log.Println("Claim[issuer] : ", claims["issuer"])
		
		} else {
			log.Println(err.Error())
			response := helper.MakeErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}