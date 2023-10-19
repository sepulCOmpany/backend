package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sepulCOmpany/backend/internal/models"
)

func (a *Api) Register(ctx *gin.Context) {
	var body models.User
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}

	dbData, err := a.dbConn.Register(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, dbData)
}

func (a *Api) Login(ctx *gin.Context) {
	var body models.User
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}

	loginedUser, err := a.dbConn.Login(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}

	ctx.JSON(http.StatusOK, loginedUser)
}
