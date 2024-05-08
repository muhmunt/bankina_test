package handler

import (
	"go-technical-test-bankina/src/entity"
	"go-technical-test-bankina/src/helper"
	"go-technical-test-bankina/src/task"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService task.Service
}

func NewTask(service task.Service) *taskHandler {
	return &taskHandler{service}
}

func (h *taskHandler) Mount(task *gin.RouterGroup) {
	task.POST("/", h.StoreTask)
	task.GET("/", h.FindTasks)
	task.GET("/:id", h.FindTaskByID)
	task.PUT("/:id", h.UpdateTask)
	task.DELETE("/:id", h.DeleteTaskByID)
}

func (h *taskHandler) StoreTask(c *gin.Context) {
	var request task.TaskRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to Store Task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)
	request.User = authUser

	newTask, err := h.taskService.StoreTask(request)

	if err != nil {
		response := helper.APIResponse("Failed to Store Task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Store Task Successfully.", http.StatusOK, "success", task.FormatTask(newTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) FindTasks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	tasks, err := h.taskService.FindTasks(offset, pageSize)
	if err != nil {
		response := helper.APIResponse("Failed get tasks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List tasks", http.StatusOK, "success", task.FormatTasks(tasks))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) FindTaskByID(c *gin.Context) {
	var request task.TaskIDRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed get task by ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getTask, err := h.taskService.FindTaskByID(request)

	if err != nil {
		response := helper.APIResponse("Failed get task by ID.", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Task Detail", http.StatusOK, "success", task.FormatTask(getTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	var requestID task.TaskIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to update task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request task.TaskRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)
	request.User = authUser

	updateTask, err := h.taskService.UpdateTask(requestID, request)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update task successfully.", http.StatusOK, "success", task.FormatTask(updateTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) DeleteTaskByID(c *gin.Context) {
	var requestID task.TaskIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.taskService.DeleteTaskByID(requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete task successfully.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
