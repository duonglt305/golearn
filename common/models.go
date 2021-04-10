package common

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `gorm:"column:created_at;sql:DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;sql:DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
