package mathpix

// https://docs.mathpix.com/?shell#query-image-results

import (
	"net/http"
	"time"
)

type ocrResultsPayload struct {
	Payload *SearchParams `in:"body=json"`
}

// SearchParams represents the query parameters for the OCR results search endpoint
type SearchParams struct {
	// Page number for pagination (starts from 1)
	Page int `json:"page,omitempty"`

	// Number of results per page
	PerPage int `json:"per_page,omitempty"`

	// Starting datetime (inclusive) for filtering results
	FromDate time.Time `json:"from_date,omitempty"`

	// Ending datetime (exclusive) for filtering results
	ToDate time.Time `json:"to_date,omitempty"`

	// Filter results by app ID
	AppID string `json:"app_id,omitempty"`

	// Filter results containing specific text in result.text
	Text string `json:"text,omitempty"`

	// Filter results containing specific text in result.text_display
	TextDisplay string `json:"text_display,omitempty"`

	// Filter results containing specific text in result.latex_styled
	LatexStyled string `json:"latex_styled,omitempty"`

	// Filter results by tags
	Tags []string `json:"tags,omitempty"`

	// Filter results containing printed text/math
	IsPrinted *bool `json:"is_printed,omitempty"`

	// Filter results containing handwritten text/math
	IsHandwritten *bool `json:"is_handwritten,omitempty"`

	// Filter results containing tables
	ContainsTable *bool `json:"contains_table,omitempty"`

	// Filter results containing chemistry diagrams
	ContainsChemistry *bool `json:"contains_chemistry,omitempty"`

	// Filter results containing diagrams
	ContainsDiagram *bool `json:"contains_diagram,omitempty"`

	// Filter results containing triangles
	ContainsTriangle *bool `json:"contains_triangle,omitempty"`
}

// OCRResultsResponse represents the top-level response from the OCR results endpoint
type OCRResultsResponse struct {
	OCRResults []OCRResult `json:"ocr_results"`

	header http.Header `json:"-"`
}

// SetHeader sets the header for the response.
func (r *OCRResultsResponse) SetHeader(header http.Header) {
	r.header = header
}

// OCRResult represents a single OCR result entry with information about the processing
type OCRResult struct {
	// ISO timestamp of recorded result information
	Timestamp string `json:"timestamp"`

	// API endpoint used for upload (eg `/v3/text`, `/v3/strokes`, ...)
	Endpoint string `json:"endpoint"`

	// Difference between timestamp and when request was received
	Duration float64 `json:"duration"`

	// Request body arguments
	RequestArgs *RequestArgs `json:"request_args"`

	// Result body for request
	Result *ResultBody `json:"result"`

	// An object of detections for each request
	Detections *Detections `json:"detections"`
}

// RequestArgs represents the original request arguments
type RequestArgs struct {
	// Tags associated with the request
	Tags []string `json:"tags,omitempty"`

	// Requested output formats
	Formats []string `json:"formats,omitempty"`
}

// ResultBody represents the OCR processing result body
type ResultBody struct {
	// Extracted text content
	Text string `json:"text"`

	// Confidence score of the OCR result
	Confidence float64 `json:"confidence"`

	// Indicates if the text is printed
	IsPrinted bool `json:"is_printed"`

	// Unique identifier for the request
	RequestID string `json:"request_id"`

	// Indicates if the text is handwritten
	IsHandwritten bool `json:"is_handwritten"`

	// Confidence rate of the OCR result
	ConfidenceRate float64 `json:"confidence_rate"`

	// Number of degrees the image was automatically rotated
	AutoRotateDegrees int `json:"auto_rotate_degrees"`

	// Confidence score for the auto-rotation
	AutoRotateConfidence float64 `json:"auto_rotate_confidence"`

	// Version of the OCR model used
	Version string `json:"version"`
}

// Detections represents various content type detection flags
type Detections struct {
	// Indicates presence of chemical formulas/diagrams
	ContainsChemistry bool `json:"contains_chemistry"`

	// Indicates presence of diagrams
	ContainsDiagram bool `json:"contains_diagram"`

	// Indicates presence of handwritten content
	IsHandwritten bool `json:"is_handwritten"`

	// Indicates presence of printed content
	IsPrinted bool `json:"is_printed"`

	// Indicates presence of tables
	ContainsTable bool `json:"contains_table"`

	// Indicates presence of triangles
	ContainsTriangle bool `json:"contains_triangle"`
}
