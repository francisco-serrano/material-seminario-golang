package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Ping(ctx *gin.Context) {
	fmt.Println("Endpoint")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	msgBefore := "Middleware BEFORE"
	msgAfter := "Middleware AFTER"

	return func(ctx *gin.Context) {
		fmt.Println(msgBefore)

		ctx.Next()

		fmt.Println(msgAfter)
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(ctx *gin.Context) {
		fmt.Println("Analyzing token")

		token := ctx.Request.FormValue("api_token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			return
		}

		ctx.Next()

		fmt.Println("After endpoint")
	}
}

func Run() {
	router := gin.Default()

	router.Use(DummyMiddleware())
	router.Use(TokenAuthMiddleware())

	router.GET("/ping", Ping)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
