package handler

import (
	"net/http"

	"github.com/OtavioSC/go-api/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", 
		"queryParamater").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error;err!=nil{
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
