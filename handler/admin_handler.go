package handler

import (
	"fmt"
	"net/http"
	"stad_projekt/helper"
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
	c.SetCookie("user_id", res.ID, 3600, "", "", false, true)
	user_id, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Iltimos qaytadan kiring",
		})
		return
	}
	c.Header("token", res.Token)

	//5. Login Succsesfule
	c.JSON(http.StatusOK, Response{
		Message: "Login succsesfule" + "User_Id : " + user_id + "   " + c.Request.UserAgent(),
	})
	// user agent ga ham boglashim kerak
}

func (h *Handler) CreateStadium(c *gin.Context) {
	var std_model models.CreateStadiumModel
	//1 .
	err := c.ShouldBindJSON(&std_model)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Malumotlar tulig' emas ....",
		})
		return
	}

	//2.
}

func (h *Handler) UploadsPictures(c *gin.Context) {
	// get file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "File type  not fount",
		})
		return
	}
	// get stadium_id
	stadium_id := c.PostForm("stadium_id")
	if stadium_id == "" {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "stadium_id not fount",
		})
	}

	file_name := file.Filename
	var name_png, name_jpg = ".png", ".jpg"
	for i := 1; i <= 4; i++ {
		if string(name_png[len(name_png)-i]) == string(file_name[len(file_name)-i]) || string(name_jpg[len(name_jpg)-i]) == string(file_name[len(file_name)-i]) {
			continue
		} else {
			c.JSON(http.StatusBadRequest, BadRequestModel{
				Message: "File type  not fount",
			})
			return
		}
	}

	user_id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusNotFound, BadRequestModel{
			Message: "User not faund",
		})
		return
	}

	rq := models.CheckStadiumPicture{
		StadiumId: stadium_id,
		User_Id:   fmt.Sprintf("%v", user_id),
	}

	f :=h.service.CheckStadiumPicture(rq)
	if !f {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "stadion_id vs user_id no same ",
		})
		return
	}
	ran_name := helper.RandomString()
	stadium_name := stadium_id

	// Upload the file to specific dst.
	dst := "./uploads/" + "./" + stadium_name + "/" + ran_name + ".png"
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "File not upload ",
		})
		return
	}
	
	c.JSON(http.StatusOK, Response{
		Message: "File muvoffiqiyatli yuklandi",
	})

	// stadion_id, err := h.service.CreatePicture(fmt.Sprintf("%v", user_id_any))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, BadRequestModel{
	// 		Message: "File not upload ",
	// 	})
	// 	return
	// }

	// fmt.Println(stadion_id)
}
func (h *Handler) CreateStadiumName(c *gin.Context) {
	var std_name models.CreateStadiumNameModel
	//1 .
	err := c.ShouldBindJSON(&std_name)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Malumotlar tulig' emas ....",
		})
		return
	}
	//2 .
	if len(std_name.Name) >= 20 {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Stadium name length increased",
		})
		return
	}

	//3 .
	user_id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusNotFound, BadRequestModel{
			Message: "User not faund",
		})
		return
	}

	var msg = models.CreateStadiumNameRequest{
		User_Id: fmt.Sprintf("%v", user_id),
		Name:    std_name.Name,
	}

	std_id, err := h.service.CreateStadiumName(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequestModel{
			Message: "Create Stadium Name error",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: std_id,
	})
}
