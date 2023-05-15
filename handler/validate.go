package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	//1 .
	token := c.Request.Header.Get("token")
	if token == "" {
		fmt.Println("Bir")
		c.AbortWithStatusJSON(http.StatusUnauthorized, BadRequestModel{
			Message: "Unauthorized",
		})
		return
	}

	token_any, err := c.Cookie("token")
	//2 .
	if err!=nil {
		fmt.Println("ikki")
		c.AbortWithStatusJSON(http.StatusUnauthorized, BadRequestModel{
			Message: "Unauthorized",
		})
		return
	}

	token_string :=fmt.Sprintf("%v",token_any)
	//3 .
	if token!=token_string {
		fmt.Println("Uch")
		c.AbortWithStatusJSON(http.StatusUnauthorized, BadRequestModel{
			Message: "Unauthorized",
		})
		return
	}

	user_id, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Iltimos qaytadan kiring",
		})
		return
	}
	c.Set("user_id",user_id)
	c.Next()
}