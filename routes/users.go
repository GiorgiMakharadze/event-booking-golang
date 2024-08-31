package routes

import (
	"net/http"

	"github.com/GiorgiMakharadze/event-booking-golang/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful!"})
}
