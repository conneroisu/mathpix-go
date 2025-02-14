package mathpix

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
