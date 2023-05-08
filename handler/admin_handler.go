package handler

import (
	"net/http"
	"stad_projekt/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var sign_model models.SignInModel
	//1.

	err := c.ShouldBindJSON(&sign_model)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Login yoki Password bush",
		})
		return
	}
	//2.

	if len(sign_model.Login) > 15 {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Login qiymati uzunligi oshib ketdi",
		})
		return
	} else if len(sign_model.Password) > 15 {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Password qiymati uzunligi oshib ketdi",
		})
		return
	}
	//3. too DB ....
	token, err := h.service.SignIn(sign_model)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Login Yoki Parol xato",
		})
		return
	}
	c.JSON(http.StatusOK, BadRequestModel{
		Message: "Asssalomu Alaykum: ..." + token,
	})
}
