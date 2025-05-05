package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Geometry represents a PostGIS geometry
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
	SRID        int         `json:"srid,omitempty"`
}

// Value implements the driver.Valuer interface for Geometry
func (g Geometry) Value() (driver.Value, error) {
	if g.Type == "" || g.Coordinates == nil {
		return nil, nil
	}

	// Convert to GeoJSON format
	geoJSON := map[string]interface{}{
		"type":        g.Type,
		"coordinates": g.Coordinates,
	}

	if g.SRID > 0 {
		geoJSON["crs"] = map[string]interface{}{
			"type": "name",
			"properties": map[string]interface{}{
				"name": fmt.Sprintf("EPSG:%d", g.SRID),
			},
		}
	}

	// Marshal to JSON
	data, err := json.Marshal(geoJSON)
	if err != nil {
		return nil, err
	}

	// Return as WKT format for PostGIS
	return fmt.Sprintf("ST_GeomFromGeoJSON('%s')", string(data)), nil
}

// Scan implements the sql.Scanner interface for Geometry
func (g *Geometry) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// Convert from WKT to GeoJSON
	// This is a simplified version, in a real implementation you would use PostGIS functions
	// to convert from WKT to GeoJSON
	var geoJSON map[string]interface{}
	if err := json.Unmarshal([]byte(value.(string)), &geoJSON); err != nil {
		return err
	}

	g.Type = geoJSON["type"].(string)
	g.Coordinates = geoJSON["coordinates"]

	if crs, ok := geoJSON["crs"].(map[string]interface{}); ok {
		if properties, ok := crs["properties"].(map[string]interface{}); ok {
			if name, ok := properties["name"].(string); ok {
				// Extract SRID from "EPSG:XXXX" format
				fmt.Sscanf(name, "EPSG:%d", &g.SRID)
			}
		}
	}

	return nil
}

// BikeRoute represents a bike route with geometry
type BikeRoute struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Geometry    Geometry  `json:"geometry"`
	Difficulty  string    `json:"difficulty"`
	Length      float64   `json:"length"`    // in kilometers
	Elevation   float64   `json:"elevation"` // in meters
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BikeStop represents a bike stop with geometry
type BikeStop struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Geometry    Geometry  `json:"geometry"`
	Type        string    `json:"type"`     // e.g., "station", "repair", "rest"
	Services    []string  `json:"services"` // e.g., ["repair", "restroom", "water"]
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
