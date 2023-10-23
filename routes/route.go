package routes

import (
	"fmt"
	"stad_projekt/config"
	"stad_projekt/handler"

	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"

	_ "stad_projekt/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteSetup(h handler.Handler, cfg config.Config) {
	w := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = append(
		config.AllowHeaders,
		`Content-Type,
		Content-Length,
		Accept-Encoding,
		X-CSRF-Token,
		Authorization,
		accept,
		origin,
		Cache-Control,
		X-Requested-With`,
	)
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	w.Use(cors.New(config))
	w.Use(MaxAllowed(100))
	w.Static("/static", "./static")

	//Auth ....
	auth := w.Group("auth")
	{
		auth.POST("/signin", h.SignIn)
	}

	// Admin ....
	admin := w.Group("admin", h.Validet)
	{
		admin.POST("/createdcountry", h.CreateCountry)
		admin.POST("/createdfield", h.CreateField)
		admin.POST("/uploadedpicture", h.CreatePicture)
		admin.GET("/getcountry", h.GetCountry)
		admin.GET("/getfield", h.GetField)
		admin.GET("/getdata", h.GetData)
		admin.GET("/getpicture", h.GetPicture)
		admin.GET("/getcountryid", h.GetFeildIdToList)
		admin.PUT("/updatecountry", h.UpdateCountry)
		admin.PUT("/updatefield", h.UpdateField)
		admin.DELETE("/deletecountry", h.DeleteCountry)
		admin.DELETE("/deletefield", h.DeleteField)
	}

	// Person ....
	person := w.Group("person")
	{
		person.GET("/getcountry", h.GetCountry)
		person.GET("/getfield", h.GetField)
		person.GET("/getdata", h.GetData)
		person.GET("/getpicture", h.GetPicture)
		person.GET("/getcountryid", h.GetFeildIdToList)
	}

	// ArdunioBoard ......
	ardunioboard := w.Group("ardunioboard")
	{
		ardunioboard.POST("/createddata", h.CreateData)
	}

	w.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	w.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
