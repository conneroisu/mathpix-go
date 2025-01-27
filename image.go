package mathpix

type (
	imageRequestPayload struct {
		Payload *ImageRequest `in:"body=json"`
	}
	// ImageRequest represents the request body for the Image endpoint.
	ImageRequest struct {
		// ImageURL is the URL of the image to be processed.
		// Optional.
		SourceURL string `json:"src,omitempty"`
		// Metadata si a map of key value pairs that will be added to the image metadata.
		// Optional.
		Metadata map[string]string `json:"metadata,omitempty"`
		// Tags are a list of tags that will be added to the image metadata.
		// Optional.
		Tags []string `json:"tags,omitempty"`
	}

	// ImageResponse represents the main response structure from the Math Recognition API.
	// It contains all information about the recognized mathematical expressions,
	// text content, and various metadata about the processed image.
	//
	// The response can include different types of data depending on the request parameters:
	// - Basic text and math recognition
	// - Line-by-line analysis (when include_line_data is true)
	// - Word-by-word analysis (when include_word_data is true)
	// - Geometric shape analysis
	// - Multiple alphabet detection
	ImageResponse struct {
		// RequestID is a unique identifier for debugging purposes
		RequestID string `json:"request_id,omitempty"`

		// Text contains the recognized text content in plain text format
		Text string `json:"text,omitempty"`

		// LatexStyled contains the mathematical expression in LaTeX format
		// This is provided when the image contains a single equation
		LatexStyled string `json:"latex_styled,omitempty"`

		// Confidence represents the estimated probability (0-1) that the entire recognition is 100% correct
		Confidence float64 `json:"confidence,omitempty"`

		// ConfidenceRate represents the estimated confidence (0-1) of output quality
		ConfidenceRate float64 `json:"confidence_rate,omitempty"`

		// LineData contains information about each line of text detected in the image
		// Only present when include_line_data is set to true in the request
		LineData []LineData `json:"line_data,omitempty"`

		// WordData contains information about individual words detected in the image
		// Only present when include_word_data is set to true in the request
		WordData []WordData `json:"word_data,omitempty"`

		// Data contains the mathematical expressions in different formats (e.g., ASCII math, LaTeX)
		Data []Data `json:"data,omitempty"`

		// HTML contains the annotated HTML output of the recognized content
		HTML string `json:"html,omitempty"`

		// DetectedAlphabets indicates which writing systems were found in the image
		DetectedAlphabets *DetectedAlphabet `json:"detected_alphabets,omitempty"`

		// IsPrinted indicates whether printed text was detected in the image
		IsPrinted bool `json:"is_printed,omitempty"`

		// IsHandwritten indicates whether handwritten content was detected in the image
		IsHandwritten bool `json:"is_handwritten,omitempty"`

		// AutoRotateConfidence represents the estimated probability (0-1) that the image needs rotation
		AutoRotateConfidence float64 `json:"auto_rotate_confidence,omitempty"`

		// GeometryData contains geometric information about detected elements in the image
		GeometryData []GeometryData `json:"geometry_data,omitempty"`

		// AutoRotateDegrees suggests the rotation angle needed to correct image orientation
		// Possible values are 0, 90, -90, 180
		AutoRotateDegrees int `json:"auto_rotate_degrees,omitempty"`

		// Error contains any error message in US locale format
		Error string `json:"error,omitempty"`

		// ErrorInfo contains detailed information about any errors that occurred
		ErrorInfo *ErrorInfo `json:"error_info,omitempty"`

		// Version is an opaque string useful for tracking differences in results
		// It changes when training data or processing methods are updated
		Version string `json:"version"`
	}
)

// Data represents mathematical expressions in different notation formats.
// Each Data object contains the expression in a specific format (e.g., ASCII math or LaTeX).
// Multiple formats can be requested for the same content.
type Data struct {
	// Type specifies the format of the mathematical expression
	// Common values: "asciimath", "latex"
	Type string `json:"type"`

	// Value contains the actual mathematical expression in the specified format
	Value string `json:"value"`
}

// LineData represents information about a detected line in the image.
// Lines can contain text, math, tables, diagrams, or other content types.
// Lines that cannot be processed or contain extraneous content will have
// included=false and may contain an error_id.
//
// The text, html, and data fields in the top-level response can be recreated
// by concatenating information from the included LineData objects.
//
// Error codes in error_id field:
// - image_not_supported: OCR engine doesn't accept this type of line
// - image_max_size: line is larger than maximal size which OCR engine supports
// - math_confidence: OCR engine failed to confidently recognize the content
// - image_no_content: line has invalid spatial dimensions (e.g., zero height)
type LineData struct {
	// Type specifies the content type of the line. Possible values:
	// "text", "math", "table", "diagram", "equation_number", "diagram_info",
	// "chart", "form_field", "code", "pseudocode", "page_info"
	Type string `json:"type"`

	// Subtype provides additional type information for specific content types:
	// - For diagrams: "chemistry", "triangle"
	// - For charts: "column", "bar", "line", "pie", "area", "scatter", "analytical"
	// - For form fields: "checkbox", "circle", "dashed"
	Subtype string `json:"subtype,omitempty"`

	// Cnt represents the contour of the line as a list of [x,y] pixel coordinates
	Cnt [][2]int `json:"cnt"`

	// Included indicates whether this line is included in the top-level OCR result
	Included bool `json:"included"`

	// IsPrinted indicates whether the line contains printed text
	IsPrinted bool `json:"is_printed"`

	// IsHandwritten indicates whether the line contains handwritten text
	IsHandwritten bool `json:"is_handwritten"`

	// ErrorID provides the reason why the line was not included in the final result
	ErrorID string `json:"error_id,omitempty"`

	// Text contains the recognized content for this line in Mathpix Markdown format
	Text string `json:"text,omitempty"`

	// Confidence represents the estimated probability (0-1) that recognition is 100% correct
	Confidence float64 `json:"confidence,omitempty"`

	// ConfidenceRate represents the estimated confidence (0-1) of output quality
	ConfidenceRate float64 `json:"confidence_rate,omitempty"`

	// AfterHyphen indicates if this line follows a line that ended with a hyphen
	AfterHyphen bool `json:"after_hyphen,omitempty"`

	// HTML contains the annotated HTML output for the line
	HTML string `json:"html,omitempty"`

	// Data contains additional data objects associated with this line
	Data []Data `json:"data,omitempty"`
}

// WordData represents information about individual word elements detected in the image.
// This provides word-level granularity for the OCR results and is only included
// when include_word_data is set to true in the request.
//
// Word-level data can be useful for:
// - Precise positioning of recognized text
// - Individual confidence scores per word
// - Distinguishing between text and math at the word level
type WordData struct {
	// Type specifies the content type of the word. Possible values:
	// "text", "math", "table", "diagram", "equation_number"
	Type string `json:"type"`

	// Subtype provides additional type information for specific content types:
	// - For diagrams: "chemistry", "triangle"
	// Only set for certain content types
	Subtype string `json:"subtype,omitempty"`

	// Cnt represents the contour of the word as a list of [x,y] pixel coordinates
	Cnt [][2]int `json:"cnt"`

	// Text contains the recognized content for this word in Mathpix Markdown format
	Text string `json:"text,omitempty"`

	// LaTeX contains the math mode LaTeX representation of the word
	// Only present for mathematical content
	LaTeX string `json:"latex,omitempty"`

	// Confidence represents the estimated probability (0-1) that recognition is 100% correct
	Confidence float64 `json:"confidence,omitempty"`

	// ConfidenceRate represents the estimated confidence (0-1) of output quality
	ConfidenceRate float64 `json:"confidence_rate,omitempty"`
}

// DetectedAlphabet indicates which writing systems were found in the processed image.
// Each field is true if any characters from that writing system were detected,
// regardless of whether they appear in the final result.
type DetectedAlphabet struct {
	// English characters detected
	English bool `json:"en"`

	// Hindi (Devanagari script) characters detected
	Hindi bool `json:"hi"`

	// Chinese characters detected
	Chinese bool `json:"zh"`

	// Japanese (Hiragana or Katakana) characters detected
	Japanese bool `json:"ja"`

	// Korean (Hangul Jamo) characters detected
	Korean bool `json:"ko"`

	// Russian characters detected
	Russian bool `json:"ru"`

	// Thai characters detected
	Thai bool `json:"th"`

	// Tamil characters detected
	Tamil bool `json:"ta"`

	// Telugu characters detected
	Telugu bool `json:"te"`

	// Gujarati characters detected
	Gujarati bool `json:"gu"`

	// Bengali characters detected
	Bengali bool `json:"bn"`

	// Vietnamese characters detected
	Vietnamese bool `json:"vi"`
}

// GeometryData represents geometric information about elements detected in the image.
// It contains details about shapes, vertices, and labels found in geometric diagrams.
// Currently, only triangle shapes are fully supported.
type GeometryData struct {
	// Position contains the pixel coordinates for this geometric element
	Position *Position `json:"position,omitempty"`

	// ShapeList contains all shapes detected in the diagram
	ShapeList []ShapeData `json:"shape_list"`

	// LabelList contains all labels associated with the geometric elements
	LabelList []LabelData `json:"label_list"`
}

// Position represents pixel coordinates in the image.
// The coordinate system starts from the top-left corner of the image.
type Position struct {
	// X coordinate, counting from top left
	X int `json:"x"`

	// Y coordinate, counting from top left
	Y int `json:"y"`
}

// ShapeData represents a geometric shape detected in the image.
// Currently supports triangle detection, with potential for additional
// shape types in future updates.
type ShapeData struct {
	// Type specifies the type of geometric shape
	// Currently only "triangle" is supported
	Type string `json:"type"`

	// VertexList contains all vertices that make up this shape
	VertexList []VertexData `json:"vertex_list"`
}

// VertexData represents a vertex in a geometric shape.
// It includes both the position of the vertex and its connections
// to other vertices in the same shape.
type VertexData struct {
	// X coordinate for the vertex, counting from top left
	X int `json:"x"`

	// Y coordinate for the vertex, counting from top left
	Y int `json:"y"`

	// EdgeList contains indices of vertices this vertex is connected to
	// Uses 0-based indexing into ShapeData.VertexList
	EdgeList []int `json:"edge_list"`
}

// LabelData represents text labels associated with geometric elements.
// Labels can include both text and mathematical expressions in LaTeX format.
type LabelData struct {
	// Position contains the pixel coordinates for this label
	Position Position `json:"position"`

	// Text contains the OCR-detected text content of the label
	Text string `json:"text"`

	// LaTeX contains the LaTeX representation of the label content
	LaTeX string `json:"latex"`

	// Confidence represents the estimated probability (0-1) that recognition is 100% correct
	Confidence float64 `json:"confidence,omitempty"`

	// ConfidenceRate represents the estimated confidence (0-1) of output quality
	ConfidenceRate float64 `json:"confidence_rate,omitempty"`
}

// ErrorInfo provides detailed information about any errors that occurred during processing.
// This includes both machine-readable codes and human-readable messages.
type ErrorInfo struct {
	// Code is a machine-readable error code
	Code string `json:"code"`

	// Message is a human-readable error description
	Message string `json:"message"`

	// Details contains any additional error-specific information
	Details any `json:"details,omitempty"`
}

// ResponseFormat represents the format of the response.
// string
type ResponseFormat string

const (
	// FormatText represents the Mathpix image format
	FormatText ResponseFormat = "text"
	// FormatHTML represents the html rendered from text via mathpix-markdown-it
	FormatHTML ResponseFormat = "html"
	// FormatData represents the data computed from text as specified in the data_options request parameter
	FormatData ResponseFormat = "data"
	// FormatLatexStyled represents the styled Latex, returned only in cases that the whole image can be reduced to a single equation
	FormatLatexStyled ResponseFormat = "latex_styled"
)

// ImageFormat represents the image format of an image
// string
type ImageFormat string

const (
	// JPEG represents JPEG image formats (*.jpeg, *.jpg, *.jpe)
	JPEG ImageFormat = "jpeg"

	// PNG represents Portable Network Graphics format (*.png)
	PNG ImageFormat = "png"

	// BMP represents Windows bitmap formats (*.bmp, *.dib)
	BMP ImageFormat = "bmp"

	// JPEG2000 represents JPEG 2000 format (*.jp2)
	JPEG2000 ImageFormat = "jp2"

	// WEBP represents WebP format (*.webp)
	WEBP ImageFormat = "webp"

	// PNM represents Portable image formats (*.pbm, *.pgm, *.ppm *.pxm, *.pnm)
	PNM ImageFormat = "pnm"

	// PFM represents PFM format (*.pfm)
	PFM ImageFormat = "pfm"

	// SUNRASTER represents Sun raster formats (*.sr, *.ras)
	SUNRASTER ImageFormat = "sunraster"

	// TIFF represents TIFF formats (*.tiff, *.tif)
	TIFF ImageFormat = "tiff"

	// OPENEXR represents OpenEXR Image format (*.exr)
	OPENEXR ImageFormat = "exr"

	// HDR represents Radiance HDR formats (*.hdr, *.pic)
	HDR ImageFormat = "hdr"

	// GDAL represents Raster and Vector geospatial data supported by GDAL
	GDAL ImageFormat = "gdal"
)

// String returns the string representation of the ImageFormat
func (f ImageFormat) String() string {
	return string(f)
}

// Extensions returns the file extensions associated with the ImageFormat
func (f ImageFormat) Extensions() []string {
	switch f {
	case JPEG:
		return []string{".jpeg", ".jpg", ".jpe"}
	case PNG:
		return []string{".png"}
	case BMP:
		return []string{".bmp", ".dib"}
	case JPEG2000:
		return []string{".jp2"}
	case WEBP:
		return []string{".webp"}
	case PNM:
		return []string{".pbm", ".pgm", ".ppm", ".pxm", ".pnm"}
	case PFM:
		return []string{".pfm"}
	case SUNRASTER:
		return []string{".sr", ".ras"}
	case TIFF:
		return []string{".tiff", ".tif"}
	case OPENEXR:
		return []string{".exr"}
	case HDR:
		return []string{".hdr", ".pic"}
	case GDAL:
		return []string{} // GDAL supports multiple formats, return empty slice
	default:
		return []string{}
	}
}

// IsValidExtension checks if the given file extension is valid for the ImageFormat
func (f ImageFormat) IsValidExtension(ext string) bool {
	for _, validExt := range f.Extensions() {
		if validExt == ext {
			return true
		}
	}
	return false
}

// ParseExtension returns the ImageFormat for a given file extension
func ParseExtension(ext string) (ImageFormat, bool) {
	for _, format := range []ImageFormat{
		JPEG, PNG, BMP, JPEG2000, WEBP, PNM,
		PFM, SUNRASTER, TIFF, OPENEXR, HDR, GDAL,
	} {
		if format.IsValidExtension(ext) {
			return format, true
		}
	}
	return "", false
}

// DataOptions represents configuration for various output formats of image data.
// All fields are optional and will use their default values if not specified.
type DataOptions struct {
	// IncludeSVG determines whether to include math SVG in HTML and data formats
	IncludeSVG bool `json:"include_svg,omitempty"`

	// IncludeTableHTML determines whether to include HTML for tables in HTML and data outputs
	IncludeTableHTML bool `json:"include_table_html,omitempty"`

	// IncludeLatex determines whether to include math mode latex in data and HTML outputs
	IncludeLatex bool `json:"include_latex,omitempty"`

	// IncludeTSV determines whether to include tab separated values (TSV) in data
	// and HTML outputs (tables only)
	IncludeTSV bool `json:"include_tsv,omitempty"`

	// IncludeAsciimath determines whether to include asciimath in data and HTML outputs
	IncludeAsciimath bool `json:"include_asciimath,omitempty"`

	// IncludeMathML determines whether to include mathml in data and HTML outputs
	IncludeMathML bool `json:"include_mathml,omitempty"`
}

// NewDataOptions creates a new DataOptions instance with default values.
// By default, all options are set to false.
func NewDataOptions() *DataOptions {
	return &DataOptions{}
}

// WithSVG sets the IncludeSVG option and returns the modified DataOptions.
func (o *DataOptions) WithSVG() *DataOptions {
	o.IncludeSVG = true
	return o
}

// WithTableHTML sets the IncludeTableHTML option and returns the modified DataOptions.
func (o *DataOptions) WithTableHTML() *DataOptions {
	o.IncludeTableHTML = true
	return o
}

// WithLatex sets the IncludeLatex option and returns the modified DataOptions.
func (o *DataOptions) WithLatex() *DataOptions {
	o.IncludeLatex = true
	return o
}

// WithTSV sets the IncludeTSV option and returns the modified DataOptions.
func (o *DataOptions) WithTSV() *DataOptions {
	o.IncludeTSV = true
	return o
}

// WithAsciimath sets the IncludeAsciimath option and returns the modified DataOptions.
func (o *DataOptions) WithAsciimath() *DataOptions {
	o.IncludeAsciimath = true
	return o
}

// WithMathML sets the IncludeMathML option and returns the modified DataOptions.
func (o *DataOptions) WithMathML() *DataOptions {
	o.IncludeMathML = true
	return o
}

// Reset sets all options back to their default values (false).
func (o *DataOptions) Reset() {
	o.IncludeSVG = false
	o.IncludeTableHTML = false
	o.IncludeLatex = false
	o.IncludeTSV = false
	o.IncludeAsciimath = false
	o.IncludeMathML = false
}

// Clone creates a deep copy of the DataOptions.
func (o *DataOptions) Clone() *DataOptions {
	return &DataOptions{
		IncludeSVG:       o.IncludeSVG,
		IncludeTableHTML: o.IncludeTableHTML,
		IncludeLatex:     o.IncludeLatex,
		IncludeTSV:       o.IncludeTSV,
		IncludeAsciimath: o.IncludeAsciimath,
		IncludeMathML:    o.IncludeMathML,
	}
}
