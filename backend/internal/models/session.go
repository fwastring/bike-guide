package models

import (
	"encoding/json"
	"time"
)

// Session represents a user session
type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	IP        string    `json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
}

// SessionData represents the data stored in Redis for a session
type SessionData struct {
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	IP        string    `json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
}

// ToSessionData converts a Session to a SessionData
func (s *Session) ToSessionData() *SessionData {
	return &SessionData{
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		ExpiresAt: s.ExpiresAt,
		IP:        s.IP,
		UserAgent: s.UserAgent,
	}
}

// FromSessionData converts a SessionData to a Session
func (s *Session) FromSessionData(id string, data *SessionData) {
	s.ID = id
	s.UserID = data.UserID
	s.CreatedAt = data.CreatedAt
	s.ExpiresAt = data.ExpiresAt
	s.IP = data.IP
	s.UserAgent = data.UserAgent
}

// MarshalJSON marshals a Session to JSON
func (s *Session) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalJSON unmarshals a Session from JSON
func (s *Session) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}
