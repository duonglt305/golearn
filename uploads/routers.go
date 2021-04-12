package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
	"golearn/users"
	"net/http"
)

func Routes(r *gin.RouterGroup) {
	upload := r.Group("uploads", users.JWTMiddleware())
	upload.GET("media", ListRootMedia)
	upload.GET("folders/:id", ListFolderAndFileById)
	upload.GET("files/:id", FileDetail)
}

func ListRootMedia(c *gin.Context) {
	db := common.NewConnection()
	u := users.GetUserContext(c)
	var media []Media
	err := db.Where(&Media{OwnerID: u.ID}).Where("parent_id IS NULL").Scopes(common.Pagination(c)).Find(&media).Error
	if err != nil {
		_ = c.Error(err)
	}
	serializer := ListMediaSerializer{
		Context: c,
		Media:   media,
	}
	c.JSON(http.StatusOK, serializer.Response())
}

func ListFolderAndFileById(c *gin.Context) {
}

func FileDetail(c *gin.Context) {

}
