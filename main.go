package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golearn/common"
	"golearn/uploads"
	"golearn/users"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_ = godotenv.Load()
	db := common.NewConnection()
	Migrate(db)
	r := gin.Default()
	r.Use(common.Handler())
	v1 := r.Group("api/v1")
	users.Routes(v1)
	uploads.Routes(v1)
	err := r.Run()
	if err != nil {
		context := &gin.Context{}
		context.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "Oops! Something went wrong.",
		})
	}
}
func Migrate(db *gorm.DB) {
	users.Migrate()
	_ = db.AutoMigrate(&uploads.Media{})
	_ = db.AutoMigrate(&uploads.Detail{})
}
