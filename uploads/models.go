package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
	"golearn/users"
)

type Type int8

const (
	Folder Type = 1
	File   Type = 2
)

type MediaItem struct {
	ID       uint        `gorm:"column:id;primaryKey;not null"`
	Name     string      `gorm:"column:name;not null"`
	Slug     string      `gorm:"column:slug;not null"`
	Type     Type        `gorm:"column:type;default:1"`
	ParentID uint        `gorm:"column:parent_id;default:0"`
	OwnerID  uint        `gorm:"column:owner_id;not null"`
	Owner    users.User  `gorm:"foreignKey:owner_id"`
	Detail   MediaDetail `gorm:"foreignKey:media_id"`
	Children []MediaItem `gorm:"foreignKey:parent_id"`
	common.Model
}

type MediaDetail struct {
	ID      uint   `gorm:"column:id;primaryKey;not null"`
	Mimes   string `gorm:"column:mimes"`
	Size    uint   `gorm:"column:size;default:0"`
	Path    string `gorm:"column:path;not null"`
	MediaID uint   `gorm:"column:media_id;not null"`
}

func ListMyMediaByParent(id uint, parent uint, c *gin.Context, media *[]MediaItem) (error, *common.Pagination) {
	db := common.NewConnection()
	p := common.NewPagination()
	err := db.Where(&MediaItem{OwnerID: id}).Where("parent_id = ?", parent).Scopes(common.Paginate(c, p)).Find(media).Error
	return err, p
}
