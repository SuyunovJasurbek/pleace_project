package handler

import (
	"fmt"
	"net/http"
	"stad_projekt/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateData(c *gin.Context) {
	var data models.AparatData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}
	//2.

	 id , err :=h.service.CreateData(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazaga yozilmadi",
		})
		return
	}

	fmt.Println(id)
	
	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Ma'lumotlar bazaga yozildi",
	})	 
}
func (h *Handler) GetData (c *gin.Context) {
	
	feild_id := c.Query("id")

	if feild_id == "" {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}

	//2.

	 res , err :=h.service.GetData(feild_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetPicture (c *gin.Context) {
	feild_id := c.Query("feild_id")

	if feild_id=="" {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar ",
		})
		return
	}
	//2.

	 res , err :=h.service.GetPicture(feild_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetFeildIdToList (c *gin.Context) {
	countryId := c.Query("id")

	//2.
	if countryId == "" {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}
	 res , err :=h.service.GetFeildIdToList(countryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}