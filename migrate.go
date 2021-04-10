package main

import (
	"github.com/joho/godotenv"
	"golearn/common"
	"golearn/users"
)

func main() {
	_ = godotenv.Load()
	common.ConnectDatabase()
	users.Migrate()
}
