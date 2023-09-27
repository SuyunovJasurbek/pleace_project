package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Validet(c *gin.Context) {
	token := c.GetHeader("token")
		
	if token == "" {
		c.JSON(400, Response{
			Succses: false,
			Message: "Token yuborilmadi",
		})
		c.Abort()
		return
	}

	//2. too DB ....
	res := h.service.Auth(token)
	if !res {
		c.JSON(401, Response{
			Succses: false,
			Message: "Token xato",
		})
		c.Abort()
		return
	}

}
