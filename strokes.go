package mathpix

import (
	"net/http"
)

// https://docs.mathpix.com/?shell#process-stroke-data

type requestStrokesPayload struct {
	// Strokes contains the handwriting stroke data
	Payload *RequestStrokes `in:"body=json"`
}

// RequestStrokes represents the request body for the v3/strokes endpoint
type RequestStrokes struct {
	// Strokes contains the handwriting stroke data.
	Strokes StrokesData `json:"strokes"`
}

// StrokesData contains the actual stroke coordinates
type StrokesData struct {
	// Strokes contains arrays of x and y coordinates representing the handwriting.
	Strokes StrokeCoordinates `json:"strokes"`
}

// StrokeCoordinates contains the x and y coordinates for each stroke
type StrokeCoordinates struct {
	// X contains arrays of x-coordinates, where each array represents one stroke.
	X [][]int `json:"x"`
	// Y contains arrays of y-coordinates, where each array represents one stroke.
	Y [][]int `json:"y"`
}

// StrokesResponse represents the response from the v3/strokes endpoint.
type StrokesResponse struct {
	// RequestID uniquely identifies the API request
	RequestID string `json:"request_id"`
	// IsPrinted indicates if the text appears to be printed
	IsPrinted bool `json:"is_printed"`
	// IsHandwritten indicates if the text appears to be handwritten
	IsHandwritten bool `json:"is_handwritten"`
	// AutoRotateConfidence indicates the confidence of the rotation detection
	AutoRotateConfidence float64 `json:"auto_rotate_confidence"`
	// AutoRotateDegrees indicates the detected rotation angle
	AutoRotateDegrees int `json:"auto_rotate_degrees"`
	// Confidence represents the estimated probability of 100% correct recognition
	Confidence float64 `json:"confidence"`
	// ConfidenceRate represents the estimated confidence of output quality
	ConfidenceRate float64 `json:"confidence_rate"`
	// LatexStyled contains the recognized LaTeX with styling
	LatexStyled string `json:"latex_styled"`
	// Text contains the recognized text with LaTeX delimiters
	Text string `json:"text"`
	// Version indicates the recognition model version used
	Version string `json:"version"`
	// HTML contains optional annotated HTML output
	HTML string `json:"html,omitempty"`

	headers http.Header `json:"-"`
}

// SetHeader sets the header of the response.
func (r *StrokesResponse) SetHeader(h http.Header) {
	r.headers = h
}
