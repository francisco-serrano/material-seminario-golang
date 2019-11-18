package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/material-seminario-golang/persistence/sample-api/services"
	"github.com/material-seminario-golang/persistence/sample-api/views"
	"io/ioutil"
	"net/http"
)

type CarController struct {
	Service *services.CarService
}

func (c *CarController) GetCar(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.Service.GetCar(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": result,
	})
}

func (c *CarController) CreateCar(ctx *gin.Context) {
	rawRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	var request views.CreateCarRequest
	if err := json.Unmarshal(rawRequest, &request); err != nil {
		panic(err)
	}

	response, err := c.Service.CreateCar(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": response,
	})
}
