package mathpix

import (
	"net/http"
)

type (
	appTokenPayload struct {
		Payload *AppTokenRequest `in:"body=json"`
	}
)

// AppTokenRequest represents the request parameters for creating a temporary app token
// from the Mathpix API endpoint POST /v3/app-tokens
type AppTokenRequest struct {
	// IncludeStrokesSessionID determines if the response should include a strokes_session_id
	// for live update drawing functionality
	IncludeStrokesSessionID bool `json:"include_strokes_session_id,omitempty"`

	// Expires specifies the duration in seconds for how long the app_token will last
	// Default: 300 seconds (5 minutes)
	// Range: 30-43200 seconds (12 hours) for regular app_token
	// Range: 30-300 seconds (5 minutes) when IncludeStrokesSessionID is true
	Expires int64 `json:"expires,omitempty"`
}

// AppTokenResponse represents the response from the Mathpix API when creating
// a temporary app token
type AppTokenResponse struct {
	// AppToken is the temporary token to be used in headers of v3/text, v3/latex,
	// or v3/strokes requests
	AppToken string `json:"app_token"`

	// StrokesSessionID is only included if requested via IncludeStrokesSessionID
	// Used for live digital ink pricing and SDKs
	StrokesSessionID string `json:"strokes_session_id,omitempty"`

	// AppTokenExpiresAt specifies when the app_token will expire in Unix time (seconds)
	AppTokenExpiresAt int64 `json:"app_token_expires_at"`

	headers http.Header `json:"-"`
}

// SetHeader sets the headers from the response.
func (r *AppTokenResponse) SetHeader(headers http.Header) {
	r.headers = headers
}

// NewAppTokenRequest creates a new AppTokenRequest with default values
func NewAppTokenRequest() *AppTokenRequest {
	return &AppTokenRequest{
		IncludeStrokesSessionID: false,
		Expires:                 300, // Default 5 minutes
	}
}

// WithStrokesSession enables the strokes session ID in the request
func (r *AppTokenRequest) WithStrokesSession() *AppTokenRequest {
	r.IncludeStrokesSessionID = true
	// Ensure expires is within valid range for strokes session
	if r.Expires > 300 {
		r.Expires = 300
	}
	return r
}

// WithExpiration sets the expiration time in seconds for the app token
// It validates the input based on whether strokes session is enabled
func (r *AppTokenRequest) WithExpiration(seconds int64) *AppTokenRequest {
	if r.IncludeStrokesSessionID {
		// For strokes session, limit to 30-300 seconds
		if seconds < 30 {
			seconds = 30
		} else if seconds > 300 {
			seconds = 300
		}
	} else {
		// For regular app token, limit to 30-43200 seconds
		if seconds < 30 {
			seconds = 30
		} else if seconds > 43200 {
			seconds = 43200
		}
	}
	r.Expires = seconds
	return r
}
