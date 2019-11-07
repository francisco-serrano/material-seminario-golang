package services

import (
	"errors"
	"github.com/material-seminario-golang/rest/sampleapi/models"
	"time"
)

type MessageStorage struct {
	counter  int64
	Messages []models.Message
}

func NewMessageStorage() *MessageStorage {
	return &MessageStorage{
		counter:  0,
		Messages: []models.Message{},
	}
}

func (m *MessageStorage) AddMessage(msg string) models.Message {
	message := models.Message{
		ID:        m.getAndAdd(),
		Text:      msg,
		Timestamp: time.Now(),
	}

	aux := append(m.Messages, message)

	m.Messages = aux

	return message
}

func (m *MessageStorage) GetMessages() []models.Message {
	return m.Messages
}

func (m *MessageStorage) GetMessage(id int64) (*models.Message, error) {
	for _, msg := range m.Messages {
		if msg.ID == id {
			return &msg, nil
		}
	}

	return nil, errors.New("message not found")
}

func (m *MessageStorage) getAndAdd() int64 {
	aux := m.counter

	m.counter++

	return aux
}
