package errors

import (
	"fmt"
	"os"
	"time"
)

type CustomError struct {
	ErrorLevel  string
	ServiceName string
	Timestamp   time.Time
	Message     string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("[%s][%s - %s] %s",
		c.ErrorLevel,
		c.ServiceName,
		c.Timestamp.String(),
		c.Message,
	)
}

func NewCustomError(level, message string) *CustomError {
	return &CustomError{
		ErrorLevel:  level,
		ServiceName: os.Getenv("SERVICE_NAME"),
		Timestamp:   time.Now(),
		Message:     message,
	}
}
