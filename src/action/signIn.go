package action

import (
	"errors"
	"net/http"

	"github.com/alekpopovic/ap-api/src/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignInDto struct {
	Email string `json:"email" binding:"required"`
}

func SignIn(c *gin.Context) {
	var dto SignInDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": error.ValidationErrors(verr)})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"user": 1})

	return
}
