package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sepulCOmpany/backend/internal/models"
)

func (a *Api) GetGrimziks(ctx *gin.Context) {
	dbData, err := a.dbConn.GetAllGrimziks()
	if err != nil {
		a.processDb500(ctx)
		return
	}
	ctx.JSON(http.StatusOK, dbData)
}

func (a *Api) CreateSepulca(ctx *gin.Context) {
	var body models.Sepulca
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}

	err := a.dbConn.CreateSepulca(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (a *Api) VaccinateSepulca(ctx *gin.Context) {
	var body models.Sepulca
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}
	err := a.dbConn.VaccinateSepulca(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}
	a.process204(ctx)
}

func (a *Api) RubberSepulca(ctx *gin.Context) {
	var body models.Sepulca
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}
	err := a.dbConn.RubberSepulca(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}
	a.process204(ctx)
}

func (a *Api) ChangeDeliveryState(ctx *gin.Context) {
	var body models.Sepulca
	if err := ctx.BindJSON(&body); err != nil {
		a.process422(ctx)
		return
	}
	err := a.dbConn.SetDeliveryState(body)
	if err != nil {
		a.processDb500(ctx)
		return
	}
	a.process204(ctx)
}

func (a *Api) GetAllSepulcas(ctx *gin.Context) {
	dbData, err := a.dbConn.GetAllSepulcas()
	if err != nil {
		a.processDb500(ctx)
		return
	}

	type resp struct {
		models.Sepulca
		Shmudrik models.UserWithoutPassword `json:"shmudrik"`
		Grimzik  models.UserWithoutPassword `json:"grimzik"`
	}
	response := make([]resp, 0, len(dbData))
	for i := range dbData {
		response = append(response, resp{
			Sepulca: dbData[i],
			Shmudrik: models.UserWithoutPassword{
				ID:       dbData[i].ShmurdikID,
				Username: dbData[i].Shmurdik,
			},
			Grimzik: models.UserWithoutPassword{
				ID:       dbData[i].GrimzikID,
				Username: dbData[i].Grimzik,
			},
		})
	}
	ctx.JSON(http.StatusOK, response)
}
