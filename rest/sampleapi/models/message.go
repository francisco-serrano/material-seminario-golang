package models

import "time"

type Message struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
