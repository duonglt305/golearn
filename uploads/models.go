package uploads

import "golearn/common"

type Folder struct {
	ID    uint   `gorm:"column:id;primaryKey;not null"`
	Name  string `gorm:"column:name;not null"`
	Slug  string `gorm:"column:slug;not null"`
	Files []File `gorm:"foreignKey:folder_id"`
	common.Model
}

type File struct {
	ID       uint    `gorm:"column:id;primaryKey;not null"`
	Name     string  `gorm:"column:name;not null"`
	Slug     string  `gorm:"column:slug;not null"`
	Mimes    string  `gorm:"column:mimes"`
	Size     float64 `gorm:"column:size"`
	FolderID uint    `gorm:"column:folder_id"`
	common.Model
}
