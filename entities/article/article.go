package article

import (
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Author  string `gorm:"type:text; not null"`
	Title  string `gorm:"type:text; not null"`
	Body string `gorm:"type:text"`
}
