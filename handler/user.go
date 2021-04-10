package handler

import (
	"bwa-startup/helper"
	"bwa-startup/user"
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
	//tangkap input user
	//map input  dari user ke struct

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"ERRORS": errors}
		response := helper.APIResponse("REGISTER ACCOUNT FAILED", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// newUser
	// user, err := h.userService.RegisterUser(input)
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("REGISTER ACCOUNT FAILED", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	// response := helper.APIResponse("Account has been registered", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
	// c.JSON(http.StatusOK, user)
}
