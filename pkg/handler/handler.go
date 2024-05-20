package handler

import (
	"net/http"
	"strconv"

	"github.com/anjush-bhargavan/password_generator/pkg/service"
	"github.com/gin-gonic/gin"
)

type PasswordHandler struct {
	Passwordservice service.PasswordService
}

func NewPasswordHandler(pu service.PasswordService) PasswordHandler {
	handler := &PasswordHandler{
		Passwordservice: pu,
	}
	return *handler
}

func (h *PasswordHandler) GeneratePassword(c *gin.Context) {
	lengthStr := c.DefaultQuery("length", "12")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid length"})
		return
	}

	password := h.Passwordservice.GeneratePassword(length)
	c.JSON(http.StatusOK, gin.H{"password": password})
}
