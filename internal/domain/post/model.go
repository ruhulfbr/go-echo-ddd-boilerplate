package post

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/user"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:text"`
	Content string `json:"content" gorm:"type:text"`
	UserID  uint
	User    user.User `gorm:"foreignkey:UserID"`
}
