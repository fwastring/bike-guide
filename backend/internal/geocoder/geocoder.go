// internal/geocoder/geocoder.go
package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// NominatimResult ... (same struct definition as before)
type ReverseNominatimResult struct {
	DisplayName string `json:"display_name"`
	Address     struct {
		Road        string `json:"road"`
		City        string `json:"city"`
		County      string `json:"county"`
		State       string `json:"state"`
		Postcode    string `json:"postcode"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		// Add other address components as needed
	} `json:"address"`
}

type NominatimResult struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

// GeocodeAddress takes an address string and returns a slice of Nominatim results.
func GeocodeAddress(address string) ([]NominatimResult, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	encodedAddress := url.QueryEscape(address)
	fullURL := fmt.Sprintf("%s?q=%s&format=jsonv2", baseURL, encodedAddress)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "YourFiberApp/1.0 (you@example.com)") // Update User-Agent

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoding service error: %d", resp.StatusCode)
	}

	var results []NominatimResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}
