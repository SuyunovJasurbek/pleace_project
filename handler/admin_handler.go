package handler

import (
	"net/http"
	"stad_projekt/helper"
	"stad_projekt/models"
	"time"

	"github.com/gin-gonic/gin"
)

// SignIn ....
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {string}  string	"ok"
// @Failure      400  {object}  string	"error"
// @Failure      404  {object}  string	"error"
// @Failure      500  {object}  string	"error"
// @Router       /auth/signin [post]
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

// ShowAccount godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {string}  string	"ok"
// @Failure      400  {string}  string	"error"
// @Failure      404  {string}  string	"error"
// @Failure      500  {string}  string	"error"
// @Router       /admin/createcountry [post]
func (h *Handler) CreateCountry(c *gin.Context) {
	var crt models.Country
	err := c.ShouldBindJSON(&crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Malumotlar tulig' emas",
		})
		return
	}

	id, err := h.service.CreateCountry(crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCountry{
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

	id, err := h.service.CreateField(crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseFeild{
		Succses: true,
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

	name := helper.RandomString(5)
	path := "static/" + field_id + "/" + name + ".jpg"
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "File not save",
		})
		return
	}

	var crt = models.Picture{
		FeildId: field_id,
		Url:     path,
	}

	id, err := h.service.CreatePicture(crt)
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

	countrys, err := h.service.GetCountry()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})

		return

	}

	c.JSON(http.StatusOK, countrys)
}

// Get Field ............
func (h *Handler) GetField(c *gin.Context) {

	countrys, err := h.service.GetField()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})

		return

	}

	c.JSON(http.StatusOK, countrys)
}

// Update Country ............
func (h *Handler) UpdateCountry(c *gin.Context) {
	var crt models.UpdateCountry
	err := c.ShouldBindJSON(&crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Malumotlar tulig' emas",
		})
		return
	}

	_, err = h.service.UpdateCountry(crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Update Succses",
	})
}

// Update Field ............
func (h *Handler) UpdateField(c *gin.Context) {
	var crt models.Feild
	err := c.ShouldBindJSON(&crt)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Malumotlar tulig' emas",
		})
		return
	}

	var crt2 = models.FeildToDB{
		Name:      crt.Name,
		CountryId: crt.CountryId,
		CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err = h.service.UpdateField(crt2)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Update Succses",
	})
}

// Delete Country ............
func (h *Handler) DeleteCountry(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Id not fount",
		})
		return
	}

	b, err := h.service.DeleteCountry(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: b,
	})
}

// Delete Field ............
func (h *Handler) DeleteField(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Id not fount",
		})
		return
	}

	_, err := h.service.DeleteField(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Biror joyida xatolik bor .",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Delete Succses",
	})
}
