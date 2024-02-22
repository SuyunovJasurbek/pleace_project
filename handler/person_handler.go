package handler

import (
	"net/http"
	"stad_projekt/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Message: "Assalomu Alaykum giybatchilar tolmanglar",
	})
}
func (h *Handler) SignUpPeraon(c *gin.Context) {
	var person models.SignInPersonModel
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}
	//2.

	err = h.service.SignUpPeraon(person)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazaga yozilmadi",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Ma'lumotlar bazaga yozildi",
	})
}
func (h *Handler) GetHumidity(c *gin.Context) {
	device_id := c.Query("device_id")
	humidity, err := h.service.GetHumidity(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	c.JSON(http.StatusOK, humidity)
}
func (h *Handler) GetTemperature(c *gin.Context) {
	device_id := c.Query("device_id")
	temperature, err := h.service.GetTemperature(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	c.JSON(http.StatusOK, temperature)
}
func (h *Handler) GetLight(c *gin.Context) {
	device_id := c.Query("device_id")
	light, err := h.service.GetLight(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	c.JSON(http.StatusOK, light)
}
func (h *Handler) GetHome(c *gin.Context) {
	id := c.Query("id")
	home, err := h.service.GetHome(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	c.JSON(http.StatusOK, home)
}
func (h *Handler) SignInPerson(c *gin.Context) {
	var parol models.Password
	err := c.ShouldBindJSON(&parol)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}
	//2.

	id, err := h.service.SignInPerson(parol.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	res := models.PasswordResponse{
		Id: id,
	}

	c.JSON(http.StatusOK, res)
}