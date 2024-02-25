package handler

import (
	"fmt"
	"net/http"
	"stad_projekt/models"

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

func (h *Handler) GetInactiveUsers(c *gin.Context) {
	//1.
	res, err := h.service.GetInactiveUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	//2.
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetActiveUsers(c *gin.Context) {
	//1.
	res, err := h.service.GetActiveUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	//2.
	c.JSON(http.StatusOK, res)
}

func (h *Handler) CreatePerson(c *gin.Context) {
	var person models.PersonCountry
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar tulig' emas",
		})
		return
	}
	//2.

	err = h.service.CreatePersonCountry(person)
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

func (h *Handler) GetActivePleaces(c *gin.Context) {
	//1.
	person_id := c.Query("person_id")
	fmt.Println(person_id, 1212)
	res, err := h.service.GetActivePleaces(person_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	//2.
	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdatePleace(c *gin.Context) {
	pleace_id := c.Query("id")
	err := h.service.UpdatePleace(pleace_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan o'zgartirilmadi",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Ma'lumotlar bazadan o'zgartirildi",
	})
}

func (h *Handler) UpdatePerson(c *gin.Context) {
	person_id := c.Query("id")
	err := h.service.UpdatePerson(person_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan o'zgartirilmadi",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Succses: true,
		Message: "Ma'lumotlar bazadan o'zgartirildi",
	})
}

func (h *Handler) GetHomeStaticData(c *gin.Context) {
	//1.
	var data = models.HomeDataStatic{
		Humidity: []models.GetHumidity{
			{
				Humidity: "20",
				Date:     "2021-09-01 12:00:00",
			},
			{
				Humidity: "30",
				Date:     "2021-09-01 13:00:00",
			},
			{
				Humidity: "12",
				Date:     "2021-09-01 14:00:00",
			},
			{
				Humidity: "45",
				Date:     "2021-09-01 15:00:00",
			},
			{
				Humidity: "18",
				Date:     "2021-09-01 16:00:00",
			},
			{
				Humidity: "23",
				Date:     "2021-09-01 17:00:00",
			},
			{
				Humidity: "31",
				Date:     "2021-09-01 18:00:00",
			},
			{
				Humidity: "56",
				Date:     "2021-09-01 19:00:00",
			},
			{
				Humidity: "23",
				Date:     "2021-09-01 20:00:00",
			},
			{
				Humidity: "12",
				Date:     "2021-09-01 21:00:00",
			},
		},
		Tempreature: []models.GetTemperature{
			{
				Temperature: "20",
				Date:        "2021-09-01 12:00:00",
			},
			{
				Temperature: "30",
				Date:        "2021-09-01 13:00:00",
			},
			{
				Temperature: "18",
				Date:        "2021-09-01 14:00:00",
			},
			{
				Temperature: "20",
				Date:        "2021-09-01 15:00:00",
			},
			{
				Temperature: "27",
				Date:        "2021-09-01 16:00:00",
			},
			{
				Temperature: "15",
				Date:        "2021-09-01 17:00:00",
			},
			{
				Temperature: "29",
				Date:        "2021-09-01 18:00:00",
			},
			{
				Temperature: "31",
				Date:        "2021-09-01 19:00:00",
			},
			{
				Temperature: "33",
				Date:        "2021-09-01 20:00:00",
			},
			{
				Temperature: "25",
				Date:        "2021-09-01 21:00:00",
			},
		},
		Light: []models.GetLight{
			{
				Light: "20",
				Date:  "2021-09-01 12:00:00",
			},
			{
				Light: "29",
				Date:  "2021-09-01 13:00:00",
			},
			{
				Light: "32",
				Date:  "2021-09-01 14:00:00",
			},
			{
				Light: "39",
				Date:  "2021-09-01 15:00:00",
			},
			{
				Light: "29",
				Date:  "2021-09-01 16:00:00",
			},
			{
				Light: "25",
				Date:  "2021-09-01 17:00:00",
			},
			{
				Light: "35",
				Date:  "2021-09-01 18:00:00",
			},
			{
				Light: "29",
				Date:  "2021-09-01 19:00:00",
			},
			{
				Light: "19",
				Date:  "2021-09-01 20:00:00",
			},
			{
				Light: "12",
				Date:  "2021-09-01 21:00:00",
			},
			
		},
	}
	//2.
	c.JSON(http.StatusOK, data)
}
