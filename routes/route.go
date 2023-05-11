package routes

import (
	"fmt"
	"stad_projekt/config"
	"stad_projekt/handler"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func RouteSetup(h handler.Handler,cfg config.Config) {
	w := gin.Default()
	w.Use(CORSMiddleware())
	admin := w.Group("admin")
	{
		admin.GET("/ping")
		admin.POST("/signin",h.SignIn)
		admin.POST("/createstadium",h.CreateStadium)
		admin.POST("createpictures",h.CreatePictures)
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