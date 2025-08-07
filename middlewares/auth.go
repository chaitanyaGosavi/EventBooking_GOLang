package middlewares

import (
	"eventsManagement/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(objContext *gin.Context) {
	token := objContext.Request.Header.Get("Authorization")
	if token == "" {
		objContext.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!!!!!!!"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		objContext.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!!!!!!!"})
		return
	}
	objContext.Set("userId", userId)
	objContext.Next()

}
