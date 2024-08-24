package routes

import (
	"net/http"
	"strconv"

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
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	event.ID = 1
	event.UserID = 1

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

	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, Try again later"})
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

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, Try again later"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event, Try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted!"})

}
