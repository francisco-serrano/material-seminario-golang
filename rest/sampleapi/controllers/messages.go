package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/material-seminario-golang/rest/sampleapi/services"
	"github.com/material-seminario-golang/rest/sampleapi/views"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MessageController struct {
	Storage *services.MessageStorage
}

func NewMessageController() *MessageController {
	return &MessageController{Storage: services.NewMessageStorage()}
}

func (m *MessageController) PostMessage(ctx *gin.Context) {
	rawRequest, err := ioutil.ReadAll(ctx.Request.Body)

	var request views.AddMessageRequest
	if err = json.Unmarshal(rawRequest, &request); err != nil {
		fmt.Println(err)

		statusCode := http.StatusBadRequest

		ctx.JSON(statusCode, views.BaseResponse{
			StatusCode: statusCode,
			Data: views.ErrorResponse{
				Description: err.Error(),
			},
		})

		return
	}

	msgCreated := m.Storage.AddMessage(request.Message)

	statusCode := http.StatusCreated

	ctx.JSON(statusCode, views.BaseResponse{
		StatusCode: statusCode,
		Data: views.AddMessageResponse{
			Msg: msgCreated,
		},
	})
}

func (m *MessageController) GetMessages(ctx *gin.Context) {
	messages := m.Storage.GetMessages()

	statusCode := http.StatusOK

	ctx.JSON(statusCode, views.BaseResponse{
		StatusCode: statusCode,
		Data: views.GetMessagesResponse{
			Messages: messages,
		},
	})
}

func (m *MessageController) GetMessage(ctx *gin.Context) {
	rawId := ctx.Param("id")
	if rawId == "" {
		fmt.Println("empty id")

		statusCode := http.StatusBadRequest

		ctx.JSON(statusCode, views.BaseResponse{
			StatusCode: statusCode,
			Data: views.ErrorResponse{
				Description: "empty id",
			},
		})

		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		fmt.Println(err)

		statusCode := http.StatusBadRequest

		ctx.JSON(statusCode, views.BaseResponse{
			StatusCode: statusCode,
			Data: views.ErrorResponse{
				Description: err.Error(),
			},
		})

		return
	}

	msg, err := m.Storage.GetMessage(int64(id))
	if err != nil {
		fmt.Println(err)

		statusCode := http.StatusBadRequest

		ctx.JSON(statusCode, views.BaseResponse{
			StatusCode: statusCode,
			Data: views.ErrorResponse{
				Description: err.Error(),
			},
		})

		return
	}

	statusCode := http.StatusOK

	ctx.JSON(statusCode, views.BaseResponse{
		StatusCode: statusCode,
		Data:       views.GetMessageResponse{
			Msg: *msg,
		},
	})
}
