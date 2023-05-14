package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PasswordReset(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": 1})
}
