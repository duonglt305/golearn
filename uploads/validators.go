package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
)

type FolderValidator struct {
	Name     string `form:"name" json:"name" binding:"required,max=50"`
	FolderID uint   `form:"folder_id" json:"folder_id" binding:"omitempty"`
}

func NewFolderValidator() *FolderValidator {
	v := &FolderValidator{}
	return v
}
func (v *FolderValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
