package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/Eef-M/EventHub/backend/config"
	"github.com/google/uuid"
)

func CreateRefreshToken(userID string) (string, error) {
	ctx := context.Background()

	oldToken, err := config.Redis.Get(ctx, fmt.Sprintf("user:%s:refresh_token", userID)).Result()
	if err == nil && oldToken != "" {
		_ = config.Redis.Del(ctx, fmt.Sprintf("refresh_token:%s", oldToken)).Err()
	}

	newToken := uuid.NewString()

	tokenKey := fmt.Sprintf("refresh_token:%s", newToken)
	userKey := fmt.Sprintf("user:%s:refresh_token", userID)
	ttl := time.Hour * 24 * 7

	err = config.Redis.Set(ctx, tokenKey, userID, ttl).Err()
	if err != nil {
		return "", err
	}

	err = config.Redis.Set(ctx, userKey, newToken, ttl).Err()
	if err != nil {
		// Clean up refresh_token if failed
		_ = config.Redis.Del(ctx, tokenKey).Err()
		return "", err
	}

	return newToken, nil
}

func ValidateRefreshToken(token string) (string, error) {
	ctx := context.Background()
	tokenKey := fmt.Sprintf("refresh_token:%s", token)

	userID, err := config.Redis.Get(ctx, tokenKey).Result()
	if err != nil {
		return "", err
	}

	userKey := fmt.Sprintf("user:%s:refresh_token", userID)
	activeToken, err := config.Redis.Get(ctx, userKey).Result()
	if err != nil || activeToken != token {
		return "", fmt.Errorf("invalid token for this user (may have been replaced)")
	}

	return userID, nil
}

func DeleteRefreshToken(token string) error {
	ctx := context.Background()

	tokenKey := fmt.Sprintf("refresh_token:%s", token)

	userID, err := config.Redis.Get(ctx, tokenKey).Result()
	if err == nil && userID != "" {
		userKey := fmt.Sprintf("user:%s:refresh_token", userID)
		_ = config.Redis.Del(ctx, userKey).Err()
	}

	return config.Redis.Del(ctx, tokenKey).Err()
}
