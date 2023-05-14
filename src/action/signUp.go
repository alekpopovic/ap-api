package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"user": 1})
}
