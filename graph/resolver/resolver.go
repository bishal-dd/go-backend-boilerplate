package resolver

import (
	"github.com/bisal-dd/go-backend-boilerplate/graph/resolver/user"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*user.UserResolver
}

func InitializeResolver(redis *redis.Client, db *gorm.DB) *Resolver {
	return &Resolver{
		UserResolver:    user.InitializeUserResolver(redis, db),
	}
}