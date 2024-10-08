package routes

import (
	"net/http"
	"strconv"

	"github.com/GiorgiMakharadze/event-booking-golang/middlewares"
	"github.com/GiorgiMakharadze/event-booking-golang/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events, Try again later"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not pasrse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, Try again later"})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {
	middlewares.Authenticate(ctx)

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events, Try again later"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not pasrse event id"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, Try again later"})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to update this event"})
		return
	}

	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event, Try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event updated!", "event": updateEvent})

}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not pasrse event id"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, Try again later"})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event, Try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted!"})

}
