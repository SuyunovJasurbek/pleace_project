package handler

import (
	"fmt"
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
	res, err := h.service.SignIn(sign_model)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Login yoki password xato",
		})
		return
	}

	//4. Set Cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", res.Token, 3600, "", "", false, true)
	c.Set("user_id", res.ID)
	user_id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Iltimos qaytadan kiring",
		})
		return
	}
	c.Header("token",res.Token)

	//5. Login Succsesfule
	c.JSON(http.StatusOK, Response{
		Message: "Login succsesfule" + "User_Id : " + fmt.Sprintf("%v", user_id)+"   "+c.Request.UserAgent(),
	})
	// user agent ga ham boglashim kerak 
}

func (h *Handler) CreateStadium(c *gin.Context) {

}
func (h *Handler) CreatePictures(c *gin.Context) {

}