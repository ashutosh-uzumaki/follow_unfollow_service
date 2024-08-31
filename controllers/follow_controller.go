package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"follow_service/models"
	"follow_service/services"
	"follow_service/utils"
	"log"
	"net/http"
)

type FollowController struct {
	Service services.FollowService
}

func (c *FollowController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	query, ok := requestBody["query"].(string)
	if !ok || query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}
	log.Println("Query:", query)

	userId, name, err := utils.ExtractUserData(query)
	if err != nil {
		http.Error(w, "Failed to extract user data", http.StatusBadRequest)
		return
	}
	user := models.User{
		UserId: userId,
		Name:   name,
	}
	log.Println("UserId:", user.UserId, "Name:", user.Name)
	if err := c.Service.CreateUser(context.Background(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"userId": user.UserId,
		"name":   user.Name,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query, ok := requestBody["query"].(string)
	if !ok || query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}
	log.Println("Query:", query)

	followerID, followeeID, err := utils.ExtractFollowData(query)
	if err != nil {
		http.Error(w, "Failed to extract follow data", http.StatusBadRequest)
		return
	}
	if followerID == "" {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}
	if followeeID == "" {
		http.Error(w, "Invalid followee ID", http.StatusBadRequest)
		return
	}
	log.Println(followeeID + " " + followerID)
	err = c.Service.FollowUser(context.Background(), followerID, followeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	result := fmt.Sprintf("You are now following %v", followeeID)
	json.NewEncoder(w).Encode(result)

}

func (c *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query, ok := requestBody["query"].(string)
	if !ok || query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}
	log.Println("Query:", query)

	followerID, followeeID, err := utils.ExtractUnFollowData(query)
	if err != nil {
		http.Error(w, "Failed to extract follow data", http.StatusBadRequest)
		return
	}
	if followerID == "" {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}
	if followeeID == "" {
		http.Error(w, "Invalid followee ID", http.StatusBadRequest)
		return
	}
	log.Println(followeeID + " " + followerID)
	err = c.Service.UnfollowUser(context.Background(), followerID, followeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	result := fmt.Sprintf("You have unfollowed %v", followeeID)
	json.NewEncoder(w).Encode(result)
}

func (c *FollowController) ListFollowers(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query, ok := requestBody["query"].(string)
	if !ok || query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}
	log.Println("Query:", query)
	userID, err := utils.ExtractFollowUserID(query)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	followers, err := c.Service.ListFollowers(context.Background(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(followers)
}

func (c *FollowController) ListFollowing(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query, ok := requestBody["query"].(string)
	if !ok || query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}
	log.Println("Query:", query)
	userID, err := utils.ExtractFolloweeUserID(query)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	following, err := c.Service.ListFollowing(context.Background(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(following)
}
