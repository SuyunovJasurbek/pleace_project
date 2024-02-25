package handler

import (
	"fmt"
	"net/http"
	"stad_projekt/models"
	"strconv"

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
func (h *Handler) GetAfislanyData(c *gin.Context) {
	device_id := c.Query("device_id")
	temperature, err := h.service.GetTemperature(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	light, err := h.service.GetLight(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	humidity, err := h.service.GetHumidity(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}

	day_date, err := h.service.GetDayDate(device_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Succses: false,
			Message: "Ma'lumotlar bazadan olinmadi",
		})
		return
	}
	fmt.Println("___________________")
	fmt.Println(day_date)
	fmt.Println("___________________")

	var day_dt []models.DayDate
	if len(day_date) != 0 {
		if len(temperature) != 0 && len(light) != 0 && len(humidity) != 0 {
			m := 0
			var day = day_date[m].Date[0:10]
			day_temp := ToInt(day_date[m].Temperature)
			day_light := ToInt(day_date[m].Light)
			day_humidity := ToInt(day_date[m].Humidity)
			for i := 0; i < len(day_date); i++ {
				if day == day_date[i].Date[0:10] {
					m = i
					day_temp += ToInt(day_date[m].Temperature)
					day_light += ToInt(day_date[m].Light)
					day_humidity += ToInt(day_date[m].Humidity)
				} else {
					day_dt = append(day_dt, models.DayDate{
						Day:         day,
						Temperature: strconv.Itoa(day_temp*len(day_date)/24),
						Light:       strconv.Itoa(day_light*len(day_date)/24),
						Humidity:    strconv.Itoa(day_humidity*len(day_date)/24),
					})
					day = day_date[i].Date[0:10]
					day_temp = ToInt(day_date[i].Temperature)
					day_light = ToInt(day_date[i].Light)
					day_humidity = ToInt(day_date[i].Humidity)
				}

			}

		}
	}

	var res = models.GetDeviceIDData{
		Humidity:    humidity,
		Tempreature: temperature,
		Light:       light,
		DayDate:     day_dt,
	}
	c.JSON(http.StatusOK, res)

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

func ToInt(s string) int {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(i)
	
}