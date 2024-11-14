package user

import (
	"context"

	"github.com/bisal-dd/go-backend-boilerplate/graph/model"
)

func (r *UserResolver) CountTotalUsers() (int64, error) {
    var totalUsers int64
    if err := r.db.Model(&model.User{}).Count(&totalUsers).Error; err != nil {
        return 0, err
    }
    return totalUsers, nil
}

func (r *UserResolver) FetchUsersFromDB(ctx context.Context, offset, limit int) ([]*model.User, error) {
    var users []*model.User
    if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
        return nil, err
    }
    
    return users, nil
}