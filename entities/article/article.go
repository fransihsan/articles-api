package article

import (
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Author  string `gorm:"type:text; not null; unique"`
	Title  string `gorm:"type:text; not null; unique"`
	Body string `gorm:"type:text"`
}
