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
	w.Use(CORSMiddleware())
	//_________________________________________________________________
	w.POST("admin/signin", h.SignIn)
	admin := w.Group("admin", handler.Validate)
	{
		admin.GET("/ping")
		admin.POST("/createstadiumname",h.CreateStadiumName)
		admin.POST("/uploadspictures", h.UploadsPictures)
	}

	person := w.Group("person")
	{
		person.GET("/ping", h.Ping)
	}
	w.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
