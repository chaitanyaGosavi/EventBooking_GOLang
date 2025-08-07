package routes

import (
	"eventsManagement/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Event Routes
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventById)

	//Protected Routes
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createNewEvent)
	authenticated.PUT("/events/:id", updateEventById)
	authenticated.DELETE("/events/:id", deleteEventById)
	authenticated.DELETE("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/register", cancelEventRegistration)

	//User Routes
	server.POST("/signup", createUser)
	server.POST("/login", loginUser)
}
