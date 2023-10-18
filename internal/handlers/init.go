package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sepulCOmpany/backend/internal/db"
)

type Api struct {
	dbConn *db.Db
}

func NewApi(db *db.Db) *Api {
	return &Api{dbConn: db}
}

func (a *Api) process422(ctx *gin.Context) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "Неккоректное тело запроса"})
}

func (a *Api) processDb500(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "DB ERROR!!!!!"})
}

func (a *Api) process204(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}
