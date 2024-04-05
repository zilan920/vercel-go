package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-vercel/app"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

// @Tags        Welcome
// @Summary     Hello User
// @Description Endpoint to Welcome user and say Hello "Name"
// @Param       name query string true "Name in the URL param"
// @Accept      json
// @Produce     json
// @Success     200 {object} object "success"
// @Failure     400 {object} object "Request Error or parameter missing"
// @Failure     404 {object} object "When user not found"
// @Failure     500 {object} object "Server Error"
// @Router      /hello/:name [GET]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello %v", c.Param("name"))
}

func Telegram(c *gin.Context) {
	client, err := app.NewTelegramClient(c)
	if err != nil {
		fmt.Println("init client error", err)
		return
	}
	client.HandleUpdate()
}
