package handler

import (
	"crowfundig/helper"
	"crowfundig/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user kestruct RegisterUserInput
	// struct diatas kita passsing sebagai paramater service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Register Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token, err :=  h.jswtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Account hass been registered", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}
