package models

type Follow struct {
	FollowerID string `json:"followerId"`
	FolloweeID string `json:"followeeId"`
}
