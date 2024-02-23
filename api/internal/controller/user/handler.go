package user

import "github.com/gin-gonic/gin"

type UserHandler struct {}

func (h *UserHandler) GetUsers(c *gin.Context) {}

func (h *UserHandler) GetUserById(c *gin.Context) {}

func (h *UserHandler) EditUser(c *gin.Context) {}

func (h *UserHandler) DeleteUser(c *gin.Context) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}