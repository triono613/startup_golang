package handler

import (
	"bwa-startup/helper"
	"bwa-startup/user"
	"fmt"
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

func (h *userHandler) Login(c *gin.Context) {
	//user masukan input email dan password
	//input ditangkap handler
	//mapping dr input user ke input struct
	//input struct parsing service
	//diservice mencari dgn bantuan repository user dgn email
	//mencocokan pasword
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errorMessage := gin.H{"ERRORS": err.Error()}
		response := helper.APIResponse("LOGIN FAILED", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	LoginUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(LoginUser, "tokentokentoken")
	response := helper.APIResponse("successfully login", http.StatusOK, "SUCCESS", formatter)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errorMessage := gin.H{"ERRORS": err.Error()}
		response := helper.APIResponse("Email cheking failed", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	emailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Email cheking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"IS_AVAILABLE": emailAvailable,
	}

	metaMessage := "email has been registered"

	if emailAvailable {
		metaMessage = "email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "SUCCESS", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")

	if err != nil {

		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "ERROR", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 10
	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "ERROR", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "ERROR", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(" avatar successfully uploaded image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
