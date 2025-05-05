package middleware

import (
	"bike-guide-api/internal/database"
	"bike-guide-api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GoogleUserInfo represents the user information returned by Google OAuth
type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

// AuthMiddleware creates a new authentication middleware
func AuthMiddleware(dbManager *database.DBManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get session ID from cookie
		sessionID := c.Cookies("session_id")
		if sessionID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: No session found",
			})
		}

		// Get session from Redis
		sessionData, err := dbManager.Redis.GetSession(sessionID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid session",
			})
		}

		// Parse session data
		var session models.Session
		if err := json.Unmarshal(sessionData, &session); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid session data",
			})
		}

		// Check if session is expired
		if time.Now().After(session.ExpiresAt) {
			// Delete expired session
			_ = dbManager.Redis.DeleteSession(sessionID)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Session expired",
			})
		}

		// Set user ID in context
		c.Locals("user_id", session.UserID)

		return c.Next()
	}
}

// GoogleAuthHandler handles Google OAuth authentication
func GoogleAuthHandler(dbManager *database.DBManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get environment variables
		clientID := os.Getenv("GOOGLE_CLIENT_ID")
		redirectURI := os.Getenv("GOOGLE_REDIRECT_URI")

		if clientID == "" || redirectURI == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Google OAuth configuration is missing",
			})
		}

		// Construct the Google OAuth URL
		authURL := fmt.Sprintf(
			"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email profile&access_type=offline",
			url.QueryEscape(clientID),
			url.QueryEscape(redirectURI),
		)

		// Redirect to Google's OAuth page
		return c.Redirect(authURL)
	}
}

// GoogleAuthCallbackHandler handles the OAuth callback from Google
func GoogleAuthCallbackHandler(dbManager *database.DBManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authorization code from the request
		code := c.Query("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Authorization code is required",
			})
		}

		// Exchange the authorization code for an access token
		accessToken, err := exchangeCodeForToken(code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to exchange code for token: %v", err),
			})
		}

		// Get user information from Google
		userInfo, err := getGoogleUserInfo(accessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get user info: %v", err),
			})
		}

		// Check if user exists in CouchDB
		user, err := findOrCreateUser(dbManager, userInfo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to find or create user: %v", err),
			})
		}

		// Create a new session
		sessionID := uuid.New().String()
		session := &models.Session{
			ID:        sessionID,
			UserID:    user.ID,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(24 * time.Hour), // 24-hour session
			IP:        c.IP(),
			UserAgent: c.Get("User-Agent"),
		}

		// Store session in Redis
		sessionData, err := json.Marshal(session)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create session",
			})
		}

		if err := dbManager.Redis.SetSession(sessionID, sessionData, 24*time.Hour); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to store session",
			})
		}

		// Set session cookie
		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Expires:  session.ExpiresAt,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		// Redirect to frontend with session
		return c.Redirect("/")
	}
}

// LogoutHandler handles user logout
func LogoutHandler(dbManager *database.DBManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get session ID from cookie
		sessionID := c.Cookies("session_id")
		if sessionID != "" {
			// Delete session from Redis
			_ = dbManager.Redis.DeleteSession(sessionID)
		}

		// Clear session cookie
		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.JSON(fiber.Map{
			"message": "Logout successful",
		})
	}
}

// exchangeCodeForToken exchanges an authorization code for an access token
func exchangeCodeForToken(code string) (string, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURI := os.Getenv("GOOGLE_REDIRECT_URI")

	// Create the request body
	body := fmt.Sprintf(
		"code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code",
		code, clientID, clientSecret, redirectURI,
	)

	// Create the request
	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	// Set the content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(respBody, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

// getGoogleUserInfo gets the user information from Google
func getGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	// Create the request
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var userInfo GoogleUserInfo
	if err := json.Unmarshal(respBody, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// findOrCreateUser finds or creates a user in CouchDB
func findOrCreateUser(dbManager *database.DBManager, userInfo *GoogleUserInfo) (*models.User, error) {
	// Try to find the user by Google ID
	var user models.User
	err := dbManager.CouchDB.GetDocument("users", userInfo.ID, &user)
	if err == nil {
		// User found, update if needed
		user.Email = userInfo.Email
		user.FirstName = userInfo.GivenName
		user.LastName = userInfo.FamilyName
		user.UpdatedAt = time.Now()

		// Save the updated user
		if err := dbManager.CouchDB.SaveDocument("users", user.ID, &user); err != nil {
			return nil, fmt.Errorf("error saving updated user: %w", err)
		}

		return &user, nil
	}

	// User not found, create a new one
	newUser := &models.User{
		ID:        userInfo.ID,
		Username:  userInfo.Email,
		Email:     userInfo.Email,
		FirstName: userInfo.GivenName,
		LastName:  userInfo.FamilyName,
		Role:      "user",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the new user
	if err := dbManager.CouchDB.SaveDocument("users", newUser.ID, newUser); err != nil {
		// If we get a 404 error, it means the database doesn't exist
		if strings.Contains(err.Error(), "404") {
			// Try to create the database
			if err := dbManager.CouchDB.CreateDatabase("users"); err != nil {
				return nil, fmt.Errorf("error creating users database: %w", err)
			}
			// Try saving the user again
			if err := dbManager.CouchDB.SaveDocument("users", newUser.ID, newUser); err != nil {
				return nil, fmt.Errorf("error saving new user after database creation: %w", err)
			}
		} else {
			return nil, fmt.Errorf("error saving new user: %w", err)
		}
	}

	return newUser, nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key" // Default secret key (not recommended for production)
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
