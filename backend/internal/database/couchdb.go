package database

import (
	"bike-guide-api/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// CouchDB represents a CouchDB database connection
type CouchDB struct {
	client   *http.Client
	baseURL  string
	username string
	password string
}

// NewCouchDB creates a new CouchDB connection
func NewCouchDB(cfg *config.Config) (*CouchDB, error) {
	client := &http.Client{}

	return &CouchDB{
		client:   client,
		baseURL:  cfg.GetCouchDBURL(),
		username: cfg.CouchDBUser,
		password: cfg.CouchDBPassword,
	}, nil
}

// CreateDatabase creates a new database in CouchDB
func (c *CouchDB) CreateDatabase(dbName string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.baseURL, dbName), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusPreconditionFailed {
		return fmt.Errorf("error creating database: %s", resp.Status)
	}

	return nil
}

// GetDocument retrieves a document from CouchDB
func (c *CouchDB) GetDocument(dbName, docID string, result interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.baseURL, dbName, docID), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting document: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error getting document: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// SaveDocument saves a document to CouchDB
func (c *CouchDB) SaveDocument(dbName, docID string, doc interface{}) error {
	jsonData, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("error marshaling document: %w", err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", c.baseURL, dbName, docID), strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error saving document: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error saving document: %s", resp.Status)
	}

	return nil
}

// DeleteDocument deletes a document from CouchDB
func (c *CouchDB) DeleteDocument(dbName, docID, rev string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s?rev=%s", c.baseURL, dbName, docID, url.QueryEscape(rev)), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting document: %s", resp.Status)
	}

	return nil
}

// QueryView queries a CouchDB view
func (c *CouchDB) QueryView(dbName, designDoc, viewName string, params map[string]string, result interface{}) error {
	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Add(key, value)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/_design/%s/_view/%s?%s",
		c.baseURL, dbName, designDoc, viewName, queryParams.Encode()), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error querying view: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error querying view: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
