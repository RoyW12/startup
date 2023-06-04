package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service)*userHandler{
	return &userHandler{userService}
}

func (h *userHandler)RegisterUser(c *gin.Context){
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":error}
		response := helper.APIResponse("Register account failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	Newuser,err := h.userService.RegisterUser(input)
	if err != nil{
		response :=helper.APIResponse("Register account failed",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	formatter := user.FormatUser(Newuser,"tokentokentoken")

	response :=helper.APIResponse("Account has been registered",http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK,response)
}