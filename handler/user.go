package handler

import (
	"campaignapi/helper"
	"campaignapi/user"
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
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		errResponse := helper.APIResponse("Register account failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errResponse := helper.APIResponse("Register account failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("User has been registered", 200, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		errResponse := helper.APIResponse("Login failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	loginUser, err := h.userService.LoginUser(input)
	if err != nil {
		errResponse := helper.APIResponse("Login failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	formatter := user.FormatUser(loginUser, "token")

	response := helper.APIResponse("Login succeed", 200, "success", formatter)
	c.JSON(http.StatusOK, response)
}
