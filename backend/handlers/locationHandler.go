package handlers

import (
	"backend/internal/geocoder"
	"backend/internal/overpass"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/go.geojson"
)


type Point struct {
	Lat float64
	Lon float64
}

// NominatimResult represents a single result from the Nominatim API.
type NominatimResult struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

// OSRMRoute represents a route from the OSRM API.  Only including the fields we need.
type OSRMRoute struct {
	Distance float64 `json:"distance"` // in meters
	Duration float64 `json:"duration"` // in seconds
	Geometry struct {
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
}

// OSRMResponse represents the entire OSRM API response.
type OSRMResponse struct {
	Routes []OSRMRoute `json:"routes"`
}

type LocationRequest struct {
	Address string `json:"address"`
	Distance string `json:"distance"`
}

func GetCoordinates(c *fiber.Ctx) error {
	var req LocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	results, err := geocoder.GeocodeAddress(req.Address)
	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to geoencode address"})
	}

	if len(results) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"latitude":  results[0].Lat,
			"longitude": results[0].Lon,
			"address":   results[0].DisplayName,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "address not found"})
	}
}

func generateNearbyPoints(startLat, startLon float64, totalDistanceKm float64, numPoints int) []Point {
	if numPoints < 2 {
		return []Point{{Lat: startLat, Lon: startLon}}
	}

	points := make([]Point, numPoints)
	points[0] = Point{Lat: startLat, Lon: startLon}

	// Earth radius in kilometers
	const earthRadiusKm = 6371
	angularStep := 2 * math.Pi / float64(numPoints)
	segmentAngularDistance := (totalDistanceKm / float64(numPoints)) / earthRadiusKm

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Keep track of the general direction (bearing)
	currentBearing := rand.Float64() * 2 * math.Pi // Start with a random initial bearing

	for i := 1; i < numPoints; i++ {
		lastPoint := points[i-1]
		lastLatRadians := toRadians(lastPoint.Lat)
		lastLonRadians := toRadians(lastPoint.Lon)

		// Introduce more significant randomness to the bearing change
		bearingChange := (rand.Float64() - 0.5) * math.Pi / 4 // Random change up to +/- 45 degrees

		// Update the current bearing
		currentBearing += angularStep + bearingChange

		// Ensure the bearing stays within 0 to 2*Pi
		currentBearing = math.Mod(currentBearing, 2*math.Pi)

		// Calculate the new latitude
		newLatRadians := math.Asin(math.Sin(lastLatRadians)*math.Cos(segmentAngularDistance) +
			math.Cos(lastLatRadians)*math.Sin(segmentAngularDistance)*math.Cos(currentBearing))

		// Calculate the new longitude
		newLonRadians := lastLonRadians + math.Atan2(math.Sin(currentBearing)*math.Sin(segmentAngularDistance)*math.Cos(lastLatRadians),
			math.Cos(segmentAngularDistance)-math.Sin(lastLatRadians)*math.Sin(newLatRadians))

		points[i] = Point{
			Lat: toDegrees(newLatRadians),
			Lon: toDegrees(newLonRadians),
		}
	}

	// The last point should ideally be close to the starting point to form a loop.
	// We can add the starting point again to explicitly close the loop.
	points = append(points, points[0])

	fmt.Println(points)
	return points
}

// toRadians converts degrees to radians.
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// toDegrees converts radians to degrees.
func toDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}



// getBikeRoute gets a bike route using the OSRM API.
func getBikeRoute(points []Point) (*OSRMResponse, error) {
	if len(points) < 2 {
		return nil, fmt.Errorf("at least two points (start and end) are required")
	}

	// Use a public OSRM instance. For production, you should run your own.
	// osrmBaseURL := "http://localhost:5000/route/v1/bike"
	osrmBaseURL := "http://osrm-backend-helm-osrm-backend:5000/route/v1/bike"

	// Construct the coordinates string
	var coords []string
	for _, p := range points {
		coords = append(coords, fmt.Sprintf("%f,%f", p.Lon, p.Lat))
	}
	coordinatesString := strings.Join(coords, ";")

	// Construct the OSRM URL
	osrmURL := fmt.Sprintf("%s/%s?geometries=geojson&overview=full&steps=true", // Added overview and steps for more details, adjust as needed
		osrmBaseURL, coordinatesString)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", osrmURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating OSRM request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing OSRM request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("OSRM service error: %s (status: %d, body: %s)", resp.Status, resp.StatusCode, string(body))
	}

	var osrmResponse OSRMResponse
	if err := json.NewDecoder(resp.Body).Decode(&osrmResponse); err != nil {
		return nil, fmt.Errorf("error decoding OSRM response: %w", err)
	}
	return &osrmResponse, nil
}


func FindRoute(c *fiber.Ctx) error {
	var req LocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	startAddress := req.Address

	results, err := geocoder.GeocodeAddress(startAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to geocode start address"})
	}
	startLon, err := strconv.ParseFloat(results[0].Lon, 10)
	startLat, err := strconv.ParseFloat(results[0].Lat, 10)
	startPoint := Point{Lon: startLon, Lat: startLat}
	var allPoints = make([]Point, 0)
	allPoints = append(allPoints, startPoint)

	distance, err := strconv.ParseInt(req.Distance, 10, 64)

	nearbyPoints := generateNearbyPoints(startLat, startLon, float64(distance), int(distance)/5)
	for _, otherPoint := range nearbyPoints {
		allPoints = append(allPoints, otherPoint)
	}
	allPoints = append(allPoints, startPoint)

	var geometry *geojson.Feature
	route, err := getBikeRoute(allPoints)
	if err == nil && route.Routes != nil && len(route.Routes) > 0 {
		for _, route := range route.Routes {
			geometry = geojson.NewLineStringFeature(route.Geometry.Coordinates)
		}
	}


	return c.JSON(geometry) // Or a more structured response
}

func Overpass(c *fiber.Ctx) error {
	var req LocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	startAddress := req.Address

	results, err := geocoder.GeocodeAddress(startAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to geocode start address"})
	}
	startLon, err := strconv.ParseFloat(results[0].Lon, 10)
	startLat, err := strconv.ParseFloat(results[0].Lat, 10)
	startPoint := Point{Lon: startLon, Lat: startLat}

	result := overpass.Query(overpass.Point(startPoint))


	return c.JSON(result) // Or a more structured response
}
