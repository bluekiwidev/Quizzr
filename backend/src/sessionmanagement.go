package main

import (
	"fmt"
	"time"
)

func storesession(sessionID string, email string) bool {
	rdb, ctx := redisinit()
	defer rdb.Close()

	err := rdb.Set(ctx, sessionID, email, 24*time.Hour).Err()
	if err != nil {
		logger.Error(fmt.Sprint("Failed to add session in Redis", err))
		return false
	}
	return true
}

func revokesession(sessionID string) bool {
	rdb, ctx := redisinit()
	defer rdb.Close()

	err := rdb.Del(ctx, sessionID).Err()
	if err != nil {
		logger.Error(fmt.Sprint("Failed to invalidate Redis session", err))
		return false
	}
	return true
}

func sessionemail(sessionID string) string {
	rdb, ctx := redisinit()
	defer rdb.Close()

	email, err := rdb.Get(ctx, sessionID).Result()
	if err != nil {
		logger.Error(fmt.Sprint("Failed to retrieve email from Redis", err))
		return ""
	}
	return email
}
