package routes

import (
	"eventsManagement/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvents(objContext *gin.Context) {
	userId := objContext.GetInt64("userId")
	id, err := strconv.ParseInt(objContext.Param("id"), 10, 64)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Could Not Fetch Event Data"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Could Not Register Event"})
		return
	}

	objContext.JSON(http.StatusCreated, gin.H{"message": "Event Registration Successful"})
}

func cancelEventRegistration(objContext *gin.Context) {
	userId := objContext.GetInt64("userId")
	eventId, err := strconv.ParseInt(objContext.Param("id"), 10, 64)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	var objEvent models.Event
	objEvent.Id = eventId
	err = objEvent.CancelRegistration(userId)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Could Not Cancel Registration"})
		return
	}

	objContext.JSON(http.StatusCreated, gin.H{"message": "Event Registration Cancelled Successfully"})
}
