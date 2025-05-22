package handlers

import (
	"backend/models"
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	"github.com/go-kivik/kivik/v3"
	"github.com/gofiber/fiber/v2"
)

var db *kivik.DB

// Initialize CouchDB connection
func InitDB(client *kivik.Client) error {
	exists, err := client.DBExists(context.TODO(), "routes")
	if err != nil {

	}
	if !exists {
		err := client.CreateDB(context.Background(), "routes")
		if err != nil {
			return err
		}
	}
    db = client.DB(context.Background(), "routes")
    return nil
}

// Create new route
func CreateRoute(c *fiber.Ctx) error {
    var newRoute models.Route
    if err := c.BodyParser(&newRoute); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Save the new route document in CouchDB
    docId, rev, err := db.CreateDoc(context.Background(), &newRoute)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create route"})
    }

    // Return the created document
    newRoute.ID = docId
    newRoute.Rev = rev
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Route created", "route": newRoute})
}

// Get all routes
func GetRoutes(c *fiber.Ctx) error {
    rows, err := db.AllDocs(context.Background(), kivik.Options{"include_docs": true})
	fmt.Print(err)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve routes"})
    }

    var routes []models.Route = make([]models.Route, 0)
	for rows.Next() {
		var route models.Route
		if err := rows.ScanDoc(&route); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// ✅ Set ID and Rev from row metadata
		route.ID = rows.ID()

		routes = append(routes, route)
	}

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"routes": routes})
}

// Get a single route by ID
func GetRoute(c *fiber.Ctx) error {
    id := c.Params("id")
    var route models.Route
    row := db.Get(context.Background(), id)

    if err := row.ScanDoc(&route); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve route"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"route": route})
}

// Update a single route by ID
func UpdateRoute(c *fiber.Ctx) error {
	id := c.Params("id")

	// Step 1: Fetch the existing document
	var existingRoute models.Route
	row := db.Get(context.Background(), id)
	if err := row.ScanDoc(&existingRoute); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Route not found"})
	}

	// Step 2: Parse the partial update from the request body
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid update payload"})
	}

	// Step 3: Apply updates to the existing document
	for key, value := range updates {
		switch key {
		case "title":
			if v, ok := value.(string); ok {
				existingRoute.Title = v
			}
		}
	}

	// Step 4: Save updated document
	rev, err := db.Put(context.Background(), id, existingRoute)
	fmt.Print(err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update route"})
	}

	existingRoute.Rev = rev

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Route updated",
		"route": existingRoute,
	})
}


// Delete a route
func DeleteRoute(c *fiber.Ctx) error {
    id := c.Params("id")
    var route models.Route
    row := db.Get(context.Background(), id)

    if err := row.ScanDoc(&route); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve route"})
    }

    // Delete the route from the database
    if _, err := db.Delete(context.Background(), route.ID, route.Rev); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete route"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Route deleted"})
}

