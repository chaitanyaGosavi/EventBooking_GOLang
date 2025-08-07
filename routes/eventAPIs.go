package routes

import (
	"eventsManagement/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllEvents(objContext *gin.Context) {
	eventList, err := models.GetAllEvents()
	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events data"})
	}
	objContext.JSON(http.StatusOK, gin.H{"message": "Events Data fetched successfully", "eventList": eventList})
}

func createNewEvent(objContext *gin.Context) {
	var objEvent models.Event

	err := objContext.ShouldBindJSON(&objEvent)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed during parsing payload"})
		return
	}

	userId := objContext.GetInt64("userId")
	objEvent.UserId = userId
	err = objEvent.Save()
	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create events data"})
	}
	objContext.JSON(http.StatusCreated, gin.H{"message": "Event Creation Successful", "event": objEvent})
}

func getEventById(objContext *gin.Context) {
	id, err := strconv.ParseInt(objContext.Param("id"), 10, 64)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event data"})
	}

	objContext.JSON(http.StatusOK, gin.H{"message": "Events Data fetched successfully", "event": event})

}

func updateEventById(objContext *gin.Context) {
	var objEvent models.Event
	err := objContext.ShouldBindJSON(&objEvent)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed during parsing payload"})
		return
	}

	id, err := strconv.ParseInt(objContext.Param("id"), 10, 64)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	userId := objContext.GetInt64("userId")
	event, err := models.GetEventByID(id)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event Id"})
		return
	}

	if event.UserId != userId {
		objContext.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
		return
	}

	rowCount, err := objEvent.UpdateEventById(id)
	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event data"})
		return
	}

	objContext.JSON(http.StatusOK, gin.H{"message": "Events Data updated successfully", "Rows Affected": rowCount})

}

func deleteEventById(objContext *gin.Context) {
	id, err := strconv.ParseInt(objContext.Param("id"), 10, 64)

	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	userId := objContext.GetInt64("userId")
	event, err := models.GetEventByID(id)
	if err != nil {
		objContext.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event Id"})
		return
	}

	if event.UserId != userId {
		objContext.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
		return
	}

	err = models.DeleteEventById(id)
	if err != nil {
		objContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event data"})
		return
	}

	objContext.JSON(http.StatusOK, gin.H{"message": "Event Data Deleted successfully"})
}
