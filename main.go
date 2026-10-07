package main

import (
	"fmt"

	handler "github.com/bootcamp-assignment/adapter/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := gin.Default()
	defer r.Run()

	handler.Handler(r)
}
