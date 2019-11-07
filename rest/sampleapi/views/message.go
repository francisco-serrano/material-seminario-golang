package views

import "github.com/material-seminario-golang/rest/sampleapi/models"

type AddMessageRequest struct {
	Message string `json:"message" example:"this is a sample message"`
}

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	Description string `json:"description" example:"invalid json"`
}

type AddMessageResponse struct {
	Msg models.Message `json:"msg"`
}

type GetMessagesResponse struct {
	Messages []models.Message `json:"messages"`
}

type GetMessageResponse struct {
	Msg models.Message `json:"message"`
}

