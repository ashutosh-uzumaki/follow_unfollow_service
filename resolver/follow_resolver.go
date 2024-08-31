package resolvers

import (
	"context"
	"errors"
	"follow_service/models"
	"follow_service/services"
)

type FollowResolver struct {
	Service services.FollowService
}

func (r *FollowResolver) FollowUser(ctx context.Context, followerID, followeeID string) (*models.Follow, error) {
	if !r.Service.UserExists(ctx, followerID) || !r.Service.UserExists(ctx, followeeID) {
		return nil, errors.New("one or both users do not exist")
	}

	err := r.Service.FollowUser(ctx, followerID, followeeID)
	if err != nil {
		return nil, err
	}
	return &models.Follow{FollowerID: followerID, FolloweeID: followeeID}, nil
}

func (r *FollowResolver) UnfollowUser(ctx context.Context, followerID, followeeID string) (bool, error) {
	if !r.Service.UserExists(ctx, followerID) || !r.Service.UserExists(ctx, followeeID) {
		return false, errors.New("one or both users do not exist")
	}

	err := r.Service.UnfollowUser(ctx, followerID, followeeID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FollowResolver) ListFollowers(ctx context.Context, userID string) ([]*models.Follow, error) {
	if !r.Service.UserExists(ctx, userID) {
		return nil, errors.New("user does not exist")
	}

	followers, err := r.Service.ListFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (r *FollowResolver) ListFollowing(ctx context.Context, userID string) ([]*models.Follow, error) {
	if !r.Service.UserExists(ctx, userID) {
		return nil, errors.New("user does not exist")
	}

	following, err := r.Service.ListFollowing(ctx, userID)
	if err != nil {
		return nil, err
	}
	return following, nil
}

func (r *FollowResolver) CreateUser(ctx context.Context, userId string, name string) (*models.User, error) {
	if r.Service.UserExists(ctx, userId) {
		return nil, errors.New("user already exists")
	}
	user := models.User{
		UserId: userId,
		Name:   name,
	}
	err := r.Service.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &models.User{
		UserId: user.UserId,
		Name:   user.Name,
	}, nil
}
