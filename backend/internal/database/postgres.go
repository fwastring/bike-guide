package database

import (
	"bike-guide-api/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PostgresDB represents a PostgreSQL database connection
type PostgresDB struct {
	DB *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	db, err := sql.Open("postgres", cfg.GetPostgresDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening postgres connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	// Enable PostGIS extension if it doesn't exist
	_, err = db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	if err != nil {
		return nil, fmt.Errorf("error enabling postgis extension: %w", err)
	}

	return &PostgresDB{
		DB: db,
	}, nil
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

// CreateGeometryTable creates a table with geometry support
func (p *PostgresDB) CreateGeometryTable(tableName, geometryColumn string, srid int) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			%s geometry(Geometry, %d),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, tableName, geometryColumn, srid)

	_, err := p.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating geometry table: %w", err)
	}

	// Create a spatial index
	indexQuery := fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS idx_%s_%s ON %s USING GIST (%s)
	`, tableName, geometryColumn, tableName, geometryColumn)

	_, err = p.DB.Exec(indexQuery)
	if err != nil {
		return fmt.Errorf("error creating spatial index: %w", err)
	}

	return nil
}
