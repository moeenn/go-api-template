package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func LoginHandler(c *gin.Context) {
	var form LoginForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "login information",
		"email":    form.Email,
		"password": form.Password,
	})
}
