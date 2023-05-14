package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/alekpopovic/ap-api/src/action"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	serveApp()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApp() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	public := router.Group("/auth")
	public.POST("/password_reset", action.PasswordReset)
	public.POST("/password_setup", action.PasswordSetup)
	public.POST("/sign_in", action.SignIn)
	public.POST("/sign_up", action.SignUp)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	router.Run(port)
	fmt.Println(fmt.Sprintf("Server running on port %s", port))
}
