package uploads

import (
	"golearn/common"
	"golearn/users"
)

type Type int

const (
	Folder Type = 1
	File   Type = 2
)

type Media struct {
	ID       uint       `gorm:"column:id;primaryKey;not null"`
	Name     string     `gorm:"column:name;not null"`
	Slug     string     `gorm:"column:slug;not null"`
	Type     int8       `gorm:"column:type;default:1"`
	ParentID uint       `gorm:"column:parent_id;default:0"`
	OwnerID  uint       `gorm:"column:owner_id;not null"`
	Owner    users.User `gorm:"foreignKey:owner_id"`
	Detail   Detail     `gorm:"foreignKey:media_id"`
	Children []Media    `gorm:"foreignKey:parent_id"`
	common.Model
}

type Detail struct {
	ID      uint   `gorm:"column:id;primaryKey;not null"`
	Mimes   string `gorm:"column:mimes"`
	Size    uint   `gorm:"column:size;default:0"`
	Path    string `gorm:"column:path;not null"`
	MediaID uint   `gorm:"column:media_id;not null"`
}
