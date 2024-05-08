package handler

import (
	"go-technical-test-bankina/src/auth"
	"go-technical-test-bankina/src/helper"
	"go-technical-test-bankina/src/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUser(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) AuthMount(user *gin.RouterGroup) {
	user.POST("/register", h.RegisterUser)
	user.POST("/login", h.LoginUser)
}

func (h *userHandler) Mount(user *gin.RouterGroup) {
	user.GET("/", h.FindUsers)
	user.GET("/:id", h.FindUserByID)
	user.PUT("/:id", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUserByID)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var request user.RegisterRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed register account.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailableEmail, err := h.userService.CheckEmailAvailable(request.Email)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isAvailableEmail {
		errorMessage := gin.H{"errors": "Email is already taken"}
		response := helper.APIResponse("Failed register account.", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(request)

	if err != nil {
		response := helper.APIResponse("Failed register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIResponse("Failed register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account registered successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var request user.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.Login(request)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := auth.GenerateToken(userData.ID)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userData, token)

	response := helper.APIResponse("Login successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FindUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	users, err := h.userService.FindUsers(offset, pageSize)
	if err != nil {
		response := helper.APIResponse("Failed get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List users", http.StatusOK, "success", user.FormatUsers(users))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FindUserByID(c *gin.Context) {
	var request user.UserIDRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed get User by ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getUser, err := h.userService.GetUserByID(request.ID)

	if err != nil {
		response := helper.APIResponse("Failed get User by ID.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("User Detail", http.StatusOK, "success", user.FormatUserDetail(getUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var requestID user.UserIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request user.UserUpdateRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateUser, err := h.userService.UpdateUser(requestID, request)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update user successfully.", http.StatusOK, "success", user.FormatUserDetail(updateUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUserByID(c *gin.Context) {
	var requestID user.UserIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.DeleteUserByID(requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete user successfully.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
