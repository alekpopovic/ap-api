package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PasswordSetup(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": 1})
}
