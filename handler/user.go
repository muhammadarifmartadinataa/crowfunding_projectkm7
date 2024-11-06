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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
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

func (h *userHandler) Login(c *gin.Context) {
	//user memasukan input (email & password)
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct kita passing ke service
	//diservice mencari dengan bantuan repository user dengan email
	//jika ketemu maka kita perlu mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//jika login berhasil maka perlu balikan loggedinUser ke dalam format json
	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Succesfuly logedin", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}
