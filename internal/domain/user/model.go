package user

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/post"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(200);"`
	Name     string `json:"name" gorm:"type:varchar(200);"`
	Password string `json:"password" gorm:"type:varchar(200);"`
	Post     []post.Post
}
