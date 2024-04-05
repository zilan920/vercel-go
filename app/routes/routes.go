package routes

import (
	"golang-vercel/app/handler"

	_ "golang-vercel/docs"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/ping", handler.Ping)
	app.POST("/telegram-webhook", handler.Telegram)
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
