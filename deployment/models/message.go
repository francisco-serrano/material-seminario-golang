package models

import "time"

type Message struct {
	ID        int64     `json:"id" example:"0"`
	Text      string    `json:"text" example:"hello how are you"`
	Timestamp time.Time `json:"timestamp" example:"2019-10-28T14:20:37.170146584-03:00"`
}
