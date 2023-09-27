package handler

import (
	"net/http"
	"stad_projekt/helper"
	"stad_projekt/models"

	"github.com/gin-gonic/gin"
)

// SignIn ....
func (h *Handler) SignIn(c *gin.Context) {
	var sign_model models.SignInModel
	//1.

	err := c.ShouldBindJSON(&sign_model)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Login yoki Password bush",
		})
		return
	}
	//2.

	if len(sign_model.Login) > 15 {
		c.JSON(http.StatusUnauthorized, Response{
			Succses: false,
			Message: "Login qiymati uzunligi oshib ketdi",
		})
		return
	} else if len(sign_model.Password) > 15 {
		c.JSON(http.StatusUnauthorized, Response{
			Succses: false,
			Message: "Password qiymati uzunligi oshib ketdi",
		})
		return
	}

	//3. too DB ....
	res, err := h.service.SignIn(sign_model)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Succses: false,
			Message: "Login yoki password xato",
		})
		return
	}

	
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Succses: false,
			Message: "Iltimos qaytadan kiring",
		})
		return
	}


	//5. Login Succsesfule
	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: res,
	})
}

// Create Country ............
func (h *Handler) CreateCountry(c *gin.Context) {
	var crt models.Country
	err := c.ShouldBindJSON(&crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Malumotlar tulig' emas",
		})
		return
	}

	id, err :=h.service.CreateCountry(crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK,ResponseCountry{
		Succses:   true,
		CountryId: id,
	})
}

// Create Field ............
func (h *Handler) CreateField(c *gin.Context) {
	var crt models.Feild
	err := c.ShouldBindJSON(&crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Malumotlar tulig' emas",
		})
		return
	}

	id, err :=h.service.CreateField(crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK,ResponseFeild{
		Succses:   true,
		FeildId: id,
	})
}

// Create Picture ............
func (h *Handler) CreatePicture(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "File type  not fount",
		})
		return
	}

	field_id := c.PostForm("field_id")
	if field_id == "" {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Field_Id not fount",
		})
		return
	}

	name :=helper.RandomString(5)
	path :="static/"+field_id+"/"+name+".jpg"
	err = c.SaveUploadedFile(file,path )
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "File not save",
		})
		return
	}

	var crt=  models.Picture{
		FeildId: field_id,
		Url:     path,
	}

	id, err :=h.service.CreatePicture(crt) 
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}


	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: id,
	})
}

// Get Country ............
func (h *Handler) GetCountry(c *gin.Context) {
	
	countrys, err :=h.service.GetCountry()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})

		return

	}

	c.JSON(http.StatusOK, countrys)
}

func (h *Handler) GetField(c *gin.Context) {
	
	countrys, err :=h.service.GetCountry()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})

		return

	}

	c.JSON(http.StatusOK, countrys)
}