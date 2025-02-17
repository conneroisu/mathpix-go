package mathpix

import "fmt"

type (
	// ErrorResponse is the error response struct.
	ErrorResponse struct {
		Error APIError `json:"error"`
	}
	// APIError is the error struct.
	APIError struct {
		ID      ErrorID `json:"id"`
		Message *string `json:"message,omitempty"`
		Detail  *string `json:"detail,omitempty"`
	}
)

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	if e.Message == nil {
		return e.ID.String()
	}
	return fmt.Sprintf(`%s: %s`, e.ID.String(), *e.Message)
}

// ErrorID is a specific error type in the system
type ErrorID string

// HTTP Error Constants
const (
	// ErrHTTPUnauthorized is invalid credentials error (HTTP 401)
	ErrHTTPUnauthorized ErrorID = "http_unauthorized"
	// ErrHTTPMaxRequests is too many requests error (HTTP 429)
	ErrHTTPMaxRequests ErrorID = "http_max_requests"
)

// JSON Error Constants
const (
	// ErrJSONSyntax is JSON syntax error
	ErrJSONSyntax ErrorID = "json_syntax"
)

// Image Error Constants
const (
	// ErrImageMissing is missing URL in request body error
	ErrImageMissing ErrorID = "image_missing"
	// ErrImageDownload is error downloading image
	ErrImageDownload ErrorID = "image_download_error"
	// ErrImageDecode is cannot decode the image data error
	ErrImageDecode ErrorID = "image_decode_error"
	// ErrImageNoContent is no content found in image error
	ErrImageNoContent ErrorID = "image_no_content"
	// ErrImageNotSupported is image not being math or text error
	ErrImageNotSupported ErrorID = "image_not_supported"
	// ErrImageMaxSize is image too large to process error
	ErrImageMaxSize ErrorID = "image_max_size"
)

// Strokes Error Constants
const (
	// ErrStrokesMissing is missing strokes in request body error
	ErrStrokesMissing ErrorID = "strokes_missing"
	// ErrStrokesSyntaxError is incorrect JSON or strokes format error
	ErrStrokesSyntaxError ErrorID = "strokes_syntax_error"
	// ErrStrokesNoContent is no content found in strokes error
	ErrStrokesNoContent ErrorID = "strokes_no_content"
)

// Options Error Constants
const (
	// ErrOptsBadCallback is bad callback field(s) error (post?, reply?, batch_id?)
	ErrOptsBadCallback ErrorID = "opts_bad_callback"
	// ErrOptsUnknownOCR is unknown ocr option(s) error
	ErrOptsUnknownOCR ErrorID = "opts_unknown_ocr"
	// ErrOptsUnknownFormat is unknown format option(s) error
	ErrOptsUnknownFormat ErrorID = "opts_unknown_format"
	// ErrOptsNumberRequired is option must be a number error
	ErrOptsNumberRequired ErrorID = "opts_number_required"
	// ErrOptsValueOutOfRange is value not in accepted range error
	ErrOptsValueOutOfRange ErrorID = "opts_value_out_of_range"
)

// PDF Error Constants
const (
	// ErrPDFEncrypted is PDF encrypted and not readable error
	ErrPDFEncrypted ErrorID = "pdf_encrypted"
	// ErrPDFUnknownID is PDF ID expired or invalid error
	ErrPDFUnknownID ErrorID = "pdf_unknown_id"
	// ErrPDFMissing is request sent without url field error
	ErrPDFMissing ErrorID = "pdf_missing"
	// ErrPDFPageLimitExceeded is PDF exceeds maximum page limit error
	ErrPDFPageLimitExceeded ErrorID = "pdf_page_limit_exceeded"
)

// Math Error Constants
const (
	// ErrMathConfidence is low confidence error
	ErrMathConfidence ErrorID = "math_confidence"
	// ErrMathSyntax is unrecognized math error
	ErrMathSyntax ErrorID = "math_syntax"
)

// System Error Constants
const (
	// ErrBatchUnknownID is unknown batch id error
	ErrBatchUnknownID ErrorID = "batch_unknown_id"
	// ErrSysException is server error
	ErrSysException ErrorID = "sys_exception"
	// ErrSysRequestTooLarge is max request size exceeded error (5mb for images and 512kb for strokes)
	ErrSysRequestTooLarge ErrorID = "sys_request_too_large"
)

// HTTPStatusCode returns the HTTP status code for the given error ID
func (e ErrorID) HTTPStatusCode() int {
	switch e {
	case ErrHTTPUnauthorized:
		return 401
	case ErrHTTPMaxRequests:
		return 429
	default:
		return 200
	}
}

// String returns the string representation of the ErrorID
func (e ErrorID) String() string {
	return string(e)
}

func isErrorID(in ErrorID) bool {
	switch in {
	case ErrHTTPUnauthorized,
		ErrHTTPMaxRequests,
		ErrJSONSyntax,
		ErrImageMissing,
		ErrImageDownload,
		ErrImageDecode,
		ErrImageNoContent,
		ErrImageNotSupported,
		ErrImageMaxSize,
		ErrStrokesMissing,
		ErrStrokesSyntaxError,
		ErrStrokesNoContent,
		ErrOptsBadCallback,
		ErrOptsUnknownOCR,
		ErrOptsUnknownFormat,
		ErrOptsNumberRequired,
		ErrOptsValueOutOfRange,
		ErrPDFEncrypted,
		ErrPDFUnknownID,
		ErrPDFMissing,
		ErrPDFPageLimitExceeded,
		ErrMathConfidence,
		ErrMathSyntax,
		ErrBatchUnknownID,
		ErrSysException,
		ErrSysRequestTooLarge:
		return true
	default:
		return false
	}
}
