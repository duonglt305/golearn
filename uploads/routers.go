package uploads

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	upload := r.Group("uploads")
	upload.GET("folders", ListRootMedia)
	upload.GET("folders/:id", ListFolderAndFileById)
	upload.GET("files/:id", FileDetail)
}

func ListRootMedia(c *gin.Context) {

}

func ListFolderAndFileById(c *gin.Context) {
}

func FileDetail(c *gin.Context) {

}
