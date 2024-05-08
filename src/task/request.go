package task

import (
	"go-technical-test-bankina/src/entity"
)

type TaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	User        entity.User
}

type TaskIDRequest struct {
	ID int `uri:"id" binding:"required"`
}
