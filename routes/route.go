package routes

import (
	"fmt"
	"stad_projekt/config"
	"stad_projekt/handler"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func RouteSetup(h handler.Handler, cfg config.Config) {
	w := gin.Default()

	w.Static("/uploads", "./uploads")

	w.POST("admin/signin", h.SignIn)
	
	admin := w.Group("admin", handler.Validate)
	{
		admin.GET("/ping")
		admin.POST("/createpace", h.CreateStadiumName)
		admin.POST("/uploadspictures", h.UploadsPictures)
	}

	person := w.Group("person")
	{
		person.GET("/ping", h.Ping)
	}
	w.Run(fmt.Sprintf(":%s", cfg.HTTPPort))

}