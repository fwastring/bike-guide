package database

import (
	"bike-guide-api/internal/config"
	"fmt"
)

// DBManager manages all database connections
type DBManager struct {
	Postgres *PostgresDB
	CouchDB  *CouchDB
	Redis    *RedisDB
}

// NewDBManager creates a new database manager with all connections
func NewDBManager(cfg *config.Config) (*DBManager, error) {
	// Initialize PostgreSQL
	postgres, err := NewPostgresDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing postgres: %w", err)
	}

	// Initialize CouchDB
	couchdb, err := NewCouchDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing couchdb: %w", err)
	}

	// Initialize Redis
	redis, err := NewRedisDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing redis: %w", err)
	}

	return &DBManager{
		Postgres: postgres,
		CouchDB:  couchdb,
		Redis:    redis,
	}, nil
}

// Close closes all database connections
func (m *DBManager) Close() error {
	var errs []error

	if err := m.Postgres.Close(); err != nil {
		errs = append(errs, fmt.Errorf("error closing postgres: %w", err))
	}

	if err := m.Redis.Close(); err != nil {
		errs = append(errs, fmt.Errorf("error closing redis: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}

	return nil
}
