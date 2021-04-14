package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/users"
)

type FolderValidator struct {
	Name      string    `form:"name" json:"name" binding:"required,max=50"`
	FolderID  uint      `form:"folder_id" json:"folder_id" binding:"omitempty"`
	MediaItem MediaItem `json:"-"`
}

func NewFolderValidator(c *gin.Context) (*FolderValidator, error) {
	u := users.GetUserContext(c)
	v := &FolderValidator{}
	if err := c.ShouldBind(v); err != nil {
		_ = c.Error(err)
		return v, err
	}
	v.MediaItem.Name = v.Name
	v.MediaItem.ParentID = v.FolderID
	v.MediaItem.Type = Folder
	v.MediaItem.OwnerID = u.ID
	return v, nil
}
