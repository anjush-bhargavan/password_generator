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
    length, err := strconv.Atoi(c.Query("length"))
    if err != nil || length <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password length"})
        return
    }

    includeUppercase := c.Query("uppercase") == "true"
    includeLowercase := c.Query("lowercase") == "true"
    includeNumbers := c.Query("numbers") == "true"
    includeSpecial := c.Query("special") == "true"

    password, err := h.Passwordservice.GeneratePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecial)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"password": password})
}
