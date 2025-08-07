package routes

import (
	"eventsManagement/models"
	"eventsManagement/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createUser(objContext *gin.Context) {
	var objUser models.User
	err := objContext.ShouldBindJSON(&objUser)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed during parsing payload"})
		return
	}

	objUser.Id = int64(time.Now().Minute())

	err = objUser.CreateNewUser()
	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed Creating User"})
		return
	}

	objContext.JSON(http.StatusCreated, gin.H{"message": "Event Creation Successful", "user": objUser})

}

func loginUser(objContext *gin.Context) {
	var objUser models.User

	err := objContext.ShouldBindJSON(&objUser)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed during parsing payload"})
		return
	}

	err = objUser.ValidateCredentials()

	if err != nil {
		objContext.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to login user"})
		return
	}

	token, err := utils.GenerateToken(objUser.Email, objUser.Id)
	if err != nil {
		objContext.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to login user"})
		return
	}

	objContext.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})
}
