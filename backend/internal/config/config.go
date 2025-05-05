package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	ServerPort string

	// PostgreSQL configuration
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	// CouchDB configuration
	CouchDBHost     string
	CouchDBPort     int
	CouchDBUser     string
	CouchDBPassword string
	CouchDBProtocol string

	// Redis configuration
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	Server struct {
		Port string
	}

	CORS struct {
		AllowOrigins string
	}

	Google struct {
		ClientID     string
		ClientSecret string
		RedirectURL  string
	}
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	} else {
		fmt.Println("Successfully loaded .env file")
	}

	// Debug: Print environment variables
	fmt.Printf("SERVER_PORT from env: %s\n", os.Getenv("SERVER_PORT"))
	fmt.Printf("GOOGLE_CLIENT_ID from env: %s\n", os.Getenv("GOOGLE_CLIENT_ID"))
	fmt.Printf("GOOGLE_REDIRECT_URI from env: %s\n", os.Getenv("GOOGLE_REDIRECT_URI"))

	config := &Config{
		// Server configuration
		ServerPort: getEnv("SERVER_PORT", "8080"),

		// PostgreSQL configuration
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnvAsInt("POSTGRES_PORT", 5432),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "bikeguide"),

		// CouchDB configuration
		CouchDBHost:     getEnv("COUCHDB_HOST", "localhost"),
		CouchDBPort:     getEnvAsInt("COUCHDB_PORT", 5984),
		CouchDBUser:     getEnv("COUCHDB_USER", "admin"),
		CouchDBPassword: getEnv("COUCHDB_PASSWORD", "password"),
		CouchDBProtocol: getEnv("COUCHDB_PROTOCOL", "http"),

		// Redis configuration
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "bikeguide"),

		Server: struct {
			Port string
		}{
			Port: getEnv("SERVER_PORT", "8080"),
		},

		CORS: struct {
			AllowOrigins string
		}{
			AllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000"),
		},

		Google: struct {
			ClientID     string
			ClientSecret string
			RedirectURL  string
		}{
			ClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			RedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		},
	}

	return config, nil
}

// GetPostgresDSN returns the PostgreSQL connection string
func (c *Config) GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost, c.PostgresPort, c.PostgresUser, c.PostgresPassword, c.PostgresDB)
}

// GetCouchDBURL returns the CouchDB URL
func (c *Config) GetCouchDBURL() string {
	return fmt.Sprintf("%s://%s:%d", c.CouchDBProtocol, c.CouchDBHost, c.CouchDBPort)
}

// GetRedisAddr returns the Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
