package repositories

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/user"
	postRepo "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories/post"
	userRepo "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories/user"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository user.Repository
	PostRepository post.Repository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: userRepo.NewUserRepository(db),
		PostRepository: postRepo.NewPostRepository(db),
	}
}
