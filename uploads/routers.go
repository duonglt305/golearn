package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/users"
	"net/http"
	"strconv"
)

func Routes(r *gin.RouterGroup) {
	upload := r.Group("uploads", users.JWTMiddleware())
	upload.GET("media", ListMedia)
	upload.GET("media/:id", ListMedia)
	upload.GET("media/files/:id", FileDetail)
}

func ListMedia(c *gin.Context) {
	var parent = 0
	id, exists := c.Params.Get("id")
	if exists {
		parent, _ = strconv.Atoi(id)
	}
	u := users.GetUserContext(c)
	var media []MediaItem
	err, p := ListMyMediaByParent(u.ID, uint(parent), c, &media)
	if err != nil {
		_ = c.Error(err)
	}
	serializer := ListMediaSerializer{
		Context:    c,
		Media:      media,
		Pagination: p,
	}
	c.JSON(http.StatusOK, serializer.Response())
}
func FileDetail(c *gin.Context) {
	id, exists := c.Params.Get("id")
	if exists {
		id, _ := strconv.Atoi(id)
		println(id)
	}
}
