package user

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/bisal-dd/go-backend-boilerplate/graph/loaders"
	"github.com/bisal-dd/go-backend-boilerplate/graph/model"
	"github.com/bisal-dd/go-backend-boilerplate/helper/contextUtil"
	"github.com/bisal-dd/go-backend-boilerplate/helper/paginationUtil"
	"github.com/bisal-dd/go-backend-boilerplate/helper/redisUtil"
)


func (r *UserResolver) Users(ctx context.Context, first *int, after *string) (*model.UserConnection, error) {
    userId, err := contextUtil.UserIdFromContext(ctx)
    if err != nil {
        return nil, err
    }
    
    offset, limit, err := paginationUtil.CalculatePagination(first, after)
	if err != nil {
		return nil, err 
	} 
    totalUsers, err := r.CountTotalUsers()
    if err != nil {
        return nil, err
    }
    users, err := r.GetCachedUsers(ctx, userId, offset, limit)
    if err != nil {
        return nil, err
    }
    if users == nil {
        users, err = r.FetchUsersFromDB(ctx, offset, limit)
        if err != nil {
            return nil, err
        }

        if err = redisUtil.CachePages(r.redis, UsersPageGroupKey, ctx, UsersKey, users, offset, limit,userId ); err != nil {
            return nil, err
        }
    }
   
    
    edges, end := Edges(offset, limit, users)
    pageInfo := PageInfo(edges, totalUsers, end, offset )
    return &model.UserConnection{
        Edges:      edges,
        PageInfo:   pageInfo,
        TotalCount: int(totalUsers),
    }, nil
}

func (r *UserResolver) User(ctx context.Context, id string) (*model.User, error) {
	loaders := loaders.For(ctx)
    user, err := loaders.UserLoader.Load(ctx, id)
    if err != nil {
        return nil, err
    }
    return user, nil
}