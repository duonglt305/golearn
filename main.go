package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golearn/common"
	"golearn/users"
	"log"
)

func main() {
	_ = godotenv.Load()
	common.ConnectDatabase()
	r := gin.Default()
	v1 := r.Group("/api/v1")
	users.Routers(v1)
	err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
