package utils

import (
	"errors"
	"regexp"
)

func ExtractUnFollowData(query string) (string, string, error) {
	re := regexp.MustCompile(`unfollowUser\(followerId: "(.*?)", followeeId: "(.*?)"\)`)
	matches := re.FindStringSubmatch(query)
	if len(matches) != 3 {
		return "", "", errors.New("failed to parse followerId and followeeId")
	}
	return matches[1], matches[2], nil
}

func ExtractFollowData(query string) (string, string, error) {
	re := regexp.MustCompile(`followUser\(followerId: "(.*?)", followeeId: "(.*?)"\)`)
	matches := re.FindStringSubmatch(query)
	if len(matches) != 3 {
		return "", "", errors.New("failed to parse followerId and followeeId")
	}
	return matches[1], matches[2], nil
}

func ExtractUserData(query string) (string, string, error) {
	re := regexp.MustCompile(`createUser\(userId: "(.*?)", name: "(.*?)"\)`)
	matches := re.FindStringSubmatch(query)
	if len(matches) != 3 {
		return "", "", errors.New("failed to parse user data")
	}
	return matches[1], matches[2], nil
}

func ExtractFollowUserID(query string) (string, error) {
	re := regexp.MustCompile(`listFollowers\(userId: "(.*?)"\)`)
	matches := re.FindStringSubmatch(query)
	if len(matches) != 2 {
		return "", errors.New("failed to get user id")
	}

	return matches[1], nil
}

func ExtractFolloweeUserID(query string) (string, error) {
	re := regexp.MustCompile(`listFollowing\(userId: "(.*?)"\)`)
	matches := re.FindStringSubmatch(query)
	if len(matches) != 2 {
		return "", errors.New("failed to get user id")
	}

	return matches[1], nil
}
