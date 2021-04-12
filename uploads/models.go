package uploads

import "golearn/common"

type Media struct {
	ID       uint        `gorm:"column:id;primaryKey;not null"`
	Name     string      `gorm:"column:name;not null"`
	Slug     string      `gorm:"column:slug;not null"`
	Type     int         `gorm:"column:type"`
	Children []Media     `gorm:"foreignKey:parent_id"`
	ParentID uint        `gorm:"parent_id"`
	Detail   MediaDetail `gorm:"foreignKey:parent_id"`
	common.Model
}
type MediaDetail struct {
	ID      uint   `gorm:"column:id;primaryKey;not null"`
	Mimes   string `gorm:"column:mimes"`
	Size    uint   `gorm:"column:size"`
	Path    string `gorm:"column:path"`
	MediaID uint   `gorm:"column:media_id"`
}
