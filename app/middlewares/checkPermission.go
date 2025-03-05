package middlewares

// import (
// 	"admin-v1/app/helpers"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func CheckPermission(c *gin.Context) {
// 	var jwt string = c.GetHeader("Authorization")

// 	var userSub helpers.UserJWTSubject = helpers.GetTokenSubject(jwt)

// 	if err := helpers.CheckJWT(jwt); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": err.Error(),
// 		})

// 		c.Abort()

// 		return
// 	}	
	
// 	c.Next()
// }