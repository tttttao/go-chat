package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goGoGo/controllers"
	"net/http"
)

func AuthJWT(c *gin.Context) {
	// Get JWT
	jwtString := c.Query("jwt")
	// Parse the token
	token, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		return []byte(controllers.SecretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "That's not even a token"})
	}
	// Check token is Valid
	if token.Valid {
		c.Next()
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "That's not even a token"})
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Timing is everything"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Couldn't handle this token"})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Couldn't handle this token"})
	}
}
