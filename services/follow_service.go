package services

import (
	"context"
	"errors"
	"follow_service/models"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowService struct {
	Driver neo4j.DriverWithContext
}

func (s *FollowService) CreateUser(ctx context.Context, user models.User) error {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	if s.UserExists(ctx, user.UserId) {
		return errors.New("user id already exists")
	}
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx,
			"CREATE (u:User {id: $userID, name: $name})",
			map[string]interface{}{
				"userID": user.UserId,
				"name":   user.Name,
			},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func (s *FollowService) UserExists(ctx context.Context, userID string) bool {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		res, err := tx.Run(ctx,
			"MATCH (u:User {id: $userID}) RETURN u LIMIT 1",
			map[string]interface{}{"userID": userID},
		)
		if err != nil {
			return nil, err
		}
		if res.Next(ctx) {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return false
	}

	return result.(bool)
}

func (s *FollowService) FollowUser(ctx context.Context, followerID, followeeID string) error {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	if !s.UserExists(ctx, followeeID) {
		return errors.New("followee does not exist")
	}
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx,
			"MATCH (follower:User {id: $followerID}), (followee:User {id: $followeeID}) "+
				"MERGE (follower)-[:FOLLOWS]->(followee)",
			map[string]interface{}{"followerID": followerID, "followeeID": followeeID},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func (s *FollowService) UnfollowUser(ctx context.Context, followerID, followeeID string) error {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx,
			"MATCH (follower:User {id: $followerID})-[r:FOLLOWS]->(followee:User {id: $followeeID}) "+
				"DELETE r",
			map[string]interface{}{"followerID": followerID, "followeeID": followeeID},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func (s *FollowService) ListFollowers(ctx context.Context, userID string) ([]*models.Follow, error) {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	followers := []*models.Follow{}
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		res, err := tx.Run(ctx,
			"MATCH (follower:User)-[:FOLLOWS]->(followee:User {id: $userID}) "+
				"RETURN follower.id AS followerId, followee.id AS followeeId",
			map[string]interface{}{"userID": userID},
		)
		if err != nil {
			return nil, err
		}

		for res.Next(ctx) {
			followers = append(followers, &models.Follow{
				FollowerID: res.Record().Values[0].(string),
				FolloweeID: res.Record().Values[1].(string),
			})
		}

		return followers, nil
	})
	if err != nil {
		return nil, err
	}

	return result.([]*models.Follow), nil
}

// ListFollowing returns a list of users that a given user is following
func (s *FollowService) ListFollowing(ctx context.Context, userID string) ([]*models.Follow, error) {
	session := s.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	following := []*models.Follow{}
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		res, err := tx.Run(ctx,
			"MATCH (follower:User {id: $userID})-[:FOLLOWS]->(followee:User) "+
				"RETURN follower.id AS followerId, followee.id AS followeeId",
			map[string]interface{}{"userID": userID},
		)
		if err != nil {
			return nil, err
		}

		for res.Next(ctx) {
			following = append(following, &models.Follow{
				FollowerID: res.Record().Values[0].(string),
				FolloweeID: res.Record().Values[1].(string),
			})
		}

		return following, nil
	})
	if err != nil {
		return nil, err
	}

	return result.([]*models.Follow), nil
}
