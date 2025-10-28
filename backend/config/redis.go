package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	// Get Redis configuration from environment variables
	redisAddr := getRedisAddr()
	redisPassword := getRedisPassword()
	redisDB := getRedisDB()

	// Create Redis client
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Test connection with ping
	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", redisAddr, err)
	}

	fmt.Printf("✓ Redis connected successfully at %s (DB: %d)\n", redisAddr, redisDB)
}

// getRedisAddr returns Redis address from environment variables
// Supports both REDIS_ADDR (backward compatibility) and REDIS_HOST+REDIS_PORT (Docker style)
func getRedisAddr() string {
	// First, check for REDIS_ADDR (backward compatibility)
	if addr := os.Getenv("REDIS_ADDR"); addr != "" {
		return addr
	}

	// If REDIS_ADDR not set, construct from REDIS_HOST and REDIS_PORT
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost" // default to localhost
		fmt.Println("Warning: REDIS_HOST not set, using default 'localhost'")
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379" // default Redis port
		fmt.Println("Warning: REDIS_PORT not set, using default '6379'")
	}

	return fmt.Sprintf("%s:%s", host, port)
}

// getRedisPassword returns Redis password from environment
func getRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD") // Empty string if not set (no password)
}

// getRedisDB returns Redis database number from environment
func getRedisDB() int {
	dbStr := os.Getenv("REDIS_DB")
	if dbStr == "" {
		return 0 // default to DB 0
	}

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Printf("Warning: Invalid REDIS_DB value '%s', using default DB 0", dbStr)
		return 0
	}

	if db < 0 || db > 15 {
		log.Printf("Warning: REDIS_DB value '%d' out of range (0-15), using default DB 0", db)
		return 0
	}

	return db
}

// CloseRedis closes the Redis connection (useful for graceful shutdown)
func CloseRedis() error {
	if Redis != nil {
		return Redis.Close()
	}
	return nil
}
