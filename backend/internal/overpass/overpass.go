package overpass

import (
	"fmt"
	"math"

	"github.com/serjvanilla/go-overpass"
)

type Point struct {
	Lat float64
	Lon float64
}

func createBbox(coordinates Point) string { // Use overpass.Point if that's your defined type
	latCenterDeg := coordinates.Lat
	lonCenterDeg := coordinates.Lon
	distanceKm := 10.0

	// Constants
	const kmPerDegLat = 111.1 // Approximate km per degree of latitude

	// Calculate latitude offset in degrees
	// This is constant as lines of latitude are roughly parallel.
	deltaLatDeg := distanceKm / kmPerDegLat

	// Calculate North and South latitude bounds
	latNorth := latCenterDeg + deltaLatDeg
	latSouth := latCenterDeg - deltaLatDeg

	// Convert central latitude to radians for the cosine function in longitude calculation
	latCenterRad := latCenterDeg * (math.Pi / 180.0)

	// Calculate km per degree of longitude at the central latitude.
	// This varies with latitude: kmPerDegLon = (pi/180) * R_earth * cos(latitude_rad)
	// A common approximation is 111.320 * cos(latitude_rad) km.
	kmPerDegLon := 111.320 * math.Cos(latCenterRad)

	var deltaLonDeg float64
	// Handle cases at or very near the poles where kmPerDegLon can be zero or extremely small.
	if math.Abs(kmPerDegLon) < 1e-9 { // Check if effectively zero
		// At the poles, "15km East/West" means you could span all longitudes.
		// Setting deltaLonDeg to 180 will make the longitude range from
		// (lonCenter - 180) to (lonCenter + 180), effectively covering all longitudes.
		deltaLonDeg = 180.0
	} else {
		deltaLonDeg = distanceKm / kmPerDegLon
	}

	// Calculate East and West longitude bounds
	lonEast := lonCenterDeg + deltaLonDeg
	lonWest := lonCenterDeg - deltaLonDeg

	// Clamp calculated latitudes to the valid range [-90, 90].
	// This is important if the central point is near a pole and adding deltaLatDeg pushes it beyond.
	latNorth = math.Min(90.0, math.Max(-90.0, latNorth))
	latSouth = math.Min(90.0, math.Max(-90.0, latSouth))

	// The Overpass API bbox format is "south,west,north,east".
	// It allows west > east for queries crossing the 180th meridian (antimeridian).
	// Therefore, we use the calculated lonWest and lonEast directly without forcing them
	// into a -180 to 180 range in a way that would break this antimeridian-crossing property.
	// For example, a box from lonWest=170 to lonEast=190 (which is -170 normalized) is valid.
	// Overpass interprets lonWest=170, lonEast=190 as 170 through 180 and -180 through -170.

	return fmt.Sprintf("%.6f,%.6f,%.6f,%.6f", latSouth, lonWest, latNorth, lonEast)
}

func Query(coordinates Point) overpass.Result {
	client := overpass.New()

	// newCoordinate := Point{Lat: coordinates.Lon, Lon: coordinates.Lat}

	bbox := createBbox(coordinates)
	// bbox := createBbox(newCoordinate)

	queryString := fmt.Sprintf(`[out:json];nwr["amenity"="toilets"](%s);>>;out geom;`, bbox)
	fmt.Println(queryString)
	result, err := client.Query(queryString)
	if err != nil {
		// This is the most important part!
		fmt.Printf("Error executing Overpass query: %v", err)
		// Depending on your application's needs, you might:
		// 1. Return an error: `return nil, err` (if your function signature allows)
		// 2. Return an empty/default result:
		return overpass.Result{} // Or whatever your library considers an empty/error result
	}
	return result
}
