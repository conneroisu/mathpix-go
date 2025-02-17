package mathpix

import (
	"net/http"
	"time"
)

type (
	// endpoint is the endpoint type representing the request and response of an api endpoint.
	endpoint[Request, Response any] struct {
		_      [0]Request
		_      [0]Response
		name   string
		method string
	}
)

var (
	imagesEndpoint = endpoint[*imageRequestPayload, *ImageResponse]{
		method: http.MethodPost,
		name:   "v3/image",
	}
	documentsEndpoint = endpoint[*documentRequestPayload, *ResponseDocument]{
		method: http.MethodPost,
		name:   "v3/pdf",
	}
	conversionStatusEndpoint = endpoint[*resultRequestPayload, *ResponseConversionResult]{
		method: http.MethodGet,
		name:   "v3/status",
	}
	batchEndpoint = endpoint[*postBatchRequestPayload, *ResponsePostBatch]{
		method: http.MethodPost,
		name:   "v3/batch",
	}
	strokesEndpoint = endpoint[*requestStrokesPayload, *StrokesResponse]{
		method: http.MethodPost,
		name:   "v3/strokes",
	}
	appTokensEndpoint = endpoint[*appTokenPayload, *AppTokenResponse]{
		method: http.MethodPost,
		name:   "v3/app-tokens",
	}
	ocrResultsEndpoint = endpoint[*ocrResultsPayload, *OCRResultsResponse]{
		method: http.MethodGet,
		name:   "v3/ocr-results",
	}
	getBatchEndpoint = endpoint[*getBatchPayload, *GetBatchResponse]{
		method: http.MethodGet,
		name:   "v3/batch",
	}
	requestUsageEndpoint = endpoint[*usagePayload, *ResponseUsage]{
		method: http.MethodPost,
		name:   "v3/usage",
	}
)

// Payloads
type (
	imageRequestPayload struct {
		Payload *ImageRequest `in:"body=json"`
	}
	postBatchRequestPayload struct {
		Payload *RequestPostBatch `in:"body=json"`
	}
	getBatchPayload struct {
		PayloadID string `json:"-"`
	}
	documentRequestPayload struct {
		Payload *RequestDocument `in:"body=json"`
	}
	resultRequestPayload struct {
		ResultRequest ResultRequest `json:"-"`
	}
	requestStrokesPayload struct {
		// Strokes contains the handwriting stroke data
		Payload *RequestStrokes `in:"body=json"`
	}
	appTokenPayload struct {
		Payload *AppTokenRequest `in:"body=json"`
	}
	ocrResultsPayload struct {
		Payload *OCRSearchRequest `in:"body=json"`
	}
)

// Requests
type (
	// RequestStrokes represents the request body for the v3/strokes endpoint
	RequestStrokes struct {
		// Strokes contains the handwriting stroke data.
		Strokes StrokesData `json:"strokes"`
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
	// RequestPostBatch is the request body for the POST /v3/batch endpoint.
	//
	// The request body may contain any /v3/latex parameters except src and must also contain a urls parameter.
	//
	// The request may also contain an additional callback parameter to receive results after all the images in the batch have been processed.
	RequestPostBatch struct {
		URLs map[string]string `json:"urls"`
		OCR  string            `json:"ocr_behavior,omitempty"`
		// Callback *Callback         `json:"callback,omitempty"`
	}
	// RequestDocument represents the request parameters for processing a PDF file or URL.
	RequestDocument struct {
		// URL is the HTTP URL where the file can be downloaded from
		URL string `json:"url,omitempty"`
		// Streaming enables streaming of PDF pages
		Streaming bool `json:"streaming,omitempty"`
		// Metadata is a key-value object for additional information
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		// AlphabetsAllowed specifies which alphabets are allowed in the output
		AlphabetsAllowed *AlphabetsAllowed `json:"alphabets_allowed,omitempty"`
		// RemoveSpaces determines whether extra white space is removed from equations
		RemoveSpaces *bool `json:"rm_spaces,omitempty"`
		// RemoveFonts determines whether font commands are removed from equations
		RemoveFonts *bool `json:"rm_fonts,omitempty"`
		// IdiomaticEqnArrays specifies whether to use aligned, gathered, or cases instead of array environment
		IdiomaticEqnArrays bool `json:"idiomatic_eqn_arrays,omitempty"`
		// IncludeEquationTags specifies whether to include equation number tags
		IncludeEquationTags bool `json:"include_equation_tags,omitempty"`
		// IncludeSmiles enables experimental chemistry diagram OCR
		IncludeSmiles *bool `json:"include_smiles,omitempty"`
		// IncludeChemistryAsImage returns image crops for chemical diagrams
		IncludeChemistryAsImage bool `json:"include_chemistry_as_image,omitempty"`
		// NumbersDefaultToMath specifies whether numbers are always math
		NumbersDefaultToMath bool `json:"numbers_default_to_math,omitempty"`
		// MathInlineDelimiters specifies begin/end inline math delimiters
		MathInlineDelimiters []string `json:"math_inline_delimiters,omitempty"`
		// MathDisplayDelimiters specifies begin/end display math delimiters
		MathDisplayDelimiters []string `json:"math_display_delimiters,omitempty"`
		// PageRanges specifies page ranges as comma-separated string
		PageRanges string `json:"page_ranges,omitempty"`
		// EnableSpellCheck enables predictive mode for English handwriting
		EnableSpellCheck bool `json:"enable_spell_check,omitempty"`
		// AutoNumberSections enables automatic section numbering
		AutoNumberSections bool `json:"auto_number_sections,omitempty"`
		// RemoveSectionNumbering removes existing section numbering
		RemoveSectionNumbering bool `json:"remove_section_numbering,omitempty"`
		// PreserveSectionNumbering keeps existing section numbering
		PreserveSectionNumbering *bool `json:"preserve_section_numbering,omitempty"`
		// EnableTablesFallback enables advanced table processing
		EnableTablesFallback bool `json:"enable_tables_fallback,omitempty"`
		// FullwidthPunctuation controls Unicode punctuation width
		FullwidthPunctuation *bool `json:"fullwidth_punctuation,omitempty"`
		// ConversionFormats specifies output formats for conversion
		ConversionFormats ConversionFormats `json:"conversion_formats"`
	}
	// ResultRequest represents the request to the result endpoint.
	ResultRequest struct {
		PDFID string `json:"pdf_id"`
	}

	// AppTokenRequest represents the request parameters for creating a temporary app token
	// from the Mathpix API endpoint POST /v3/app-tokens
	AppTokenRequest struct {
		// IncludeStrokesSessionID determines if the response should include a strokes_session_id
		// for live update drawing functionality
		IncludeStrokesSessionID bool `json:"include_strokes_session_id,omitempty"`

		// Expires specifies the duration in seconds for how long the app_token will last
		// Default: 300 seconds (5 minutes)
		// Range: 30-43200 seconds (12 hours) for regular app_token
		// Range: 30-300 seconds (5 minutes) when IncludeStrokesSessionID is true
		Expires int64 `json:"expires,omitempty"`
	}

	// OCRSearchRequest represents the query parameters for the OCR results search endpoint
	OCRSearchRequest struct {
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
)

// Responses
type (
	// StrokesResponse represents the response from the v3/strokes endpoint.
	StrokesResponse struct {
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
	}

	// OCRResultsResponse represents the top-level response from the OCR results endpoint
	OCRResultsResponse struct {
		OCRResults []OCRResult `json:"ocr_results"`
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
	// ResponsePostBatch is the response from the batch endpoint.
	//
	// The response contains only a unique batch_id value.
	// Even if the request includes a callback, there is no guarantee the callback will run successfully (because of a transient network failure, for example).
	// The preferred approach is to wait an appropriate length of time (about one second for every five images in the batch) and then do a GET on /v3/batch/:id where :id is the batch_id value.
	// The GET request must contain the same app_id and app_key headers as the POST to /v3/batch.
	ResponsePostBatch struct {
		BatchID string `json:"batch_id"`
	}
	// GetBatchResponse is the response from the GET /v3/batch/:id endpoint.
	GetBatchResponse struct {
		Keys    []string               `json:"keys"`
		Results map[string]interface{} `json:"results"`
	}
	// ResponseDocument represents the response from the PDF processing endpoint.
	ResponseDocument struct {
		// PDFID is the tracking ID to get status and result
		PDFID string `json:"pdf_id"`
		// Error contains US locale error message if present
		Error string `json:"error,omitempty"`
		// ErrorInfo contains detailed error information
		ErrorInfo map[string]interface{} `json:"error_info,omitempty"`
	}
	// ResponseConversionResult represents the response from the result endpoint.
	ResponseConversionResult struct {
		Status     ConversionStatusType                   `json:"status"`
		Coversions map[ConversionFormats]ConversionStatus `json:"conversion_status"`
	}

	// AppTokenResponse represents the response from the Mathpix API when creating
	// a temporary app token
	AppTokenResponse struct {
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
)

// Supporting Data Types
const (
	// FormatText represents the Mathpix image format
	FormatText ResponseFormat = "text"
	// FormatHTML represents the html rendered from text via mathpix-markdown-it
	FormatHTML ResponseFormat = "html"
	// FormatData represents the data computed from text as specified in the data_options request parameter
	FormatData ResponseFormat = "data"
	// FormatLatexStyled represents the styled Latex, returned only in cases that the whole image can be reduced to a single equation
	FormatLatexStyled ResponseFormat = "latex_styled"
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
	// ConversionStatusProcessing represents a document conversion that is in progress.
	ConversionStatusProcessing ConversionStatusType = "processing"
	// ConversionStatusCompleted represents a document conversion that is completed.
	ConversionStatusCompleted ConversionStatusType = "completed"
	// ConversionStatusError represents a document conversion that has failed.
	ConversionStatusError ConversionStatusType = "error"
	// InputFormatPDF represents PDF document format
	InputFormatPDF InputFormat = "pdf"
	// InputFormatEPUB represents EPUB ebook format
	InputFormatEPUB InputFormat = "epub"
	// InputFormatDOCX represents Microsoft Word DOCX format
	InputFormatDOCX InputFormat = "docx"
	// InputFormatPPTX represents Microsoft PowerPoint PPTX format
	InputFormatPPTX InputFormat = "pptx"
	// InputFormatAZW represents Amazon Kindle AZW format
	InputFormatAZW InputFormat = "azw"
	// InputFormatAZW3 represents Amazon Kindle AZW3 format
	InputFormatAZW3 InputFormat = "azw3"
	// InputFormatKFX represents Amazon Kindle KFX format
	InputFormatKFX InputFormat = "kfx"
	// InputFormatMOBI represents Mobipocket ebook format
	InputFormatMOBI InputFormat = "mobi"
	// InputFormatDJVU represents DjVu document format
	InputFormatDJVU InputFormat = "djvu"
	// InputFormatDOC represents Microsoft Word DOC format
	InputFormatDOC InputFormat = "doc"
	// InputFormatWPD represents WordPerfect Document format
	InputFormatWPD InputFormat = "wpd"
	// InputFormatODT represents OpenDocument Text format
	InputFormatODT InputFormat = "odt"
	// DocumentFormatMMD represents Mathpix Markdown specification format.
	DocumentFormatMMD DocumentOutputFormat = "mmd"
	// DocumentFormatMD represents standard Markdown specification format.
	DocumentFormatMD DocumentOutputFormat = "md"
	// DocumentFormatDOCX represents Microsoft Word DOCX format>
	DocumentFormatDOCX DocumentOutputFormat = "docx"
	// DocumentFormatLaTeXZip represents LaTeX format with included images in ZIP.
	DocumentFormatLaTeXZip DocumentOutputFormat = "latex_zip"
	// DocumentFormatHTML represents rendered Mathpix Markdown content in HTML.
	DocumentFormatHTML DocumentOutputFormat = "html"
	// DocumentFormatPDFWithHTML represents PDF with HTML rendering.
	DocumentFormatPDFWithHTML DocumentOutputFormat = "pdf_html"
	// DocumentFormatPDFWithLaTeX represents PDF with selectable LaTeX equations.
	DocumentFormatPDFWithLaTeX DocumentOutputFormat = "pdf_latex"
)

type (
	// ResponseFormat represents the format of the response.
	// string
	ResponseFormat string
	// ImageFormat represents the image format of an image
	// string
	ImageFormat string
	// ConversionStatusType represents the status of a document conversion.
	// string
	ConversionStatusType string
	// InputFormat represents supported input file formats for Mathpix processing.
	// It is used to specify the format of documents being submitted for processing.
	InputFormat string
	// DocumentOutputFormat represents supported output file formats from Mathpix processing.
	// It defines the format in which processed documents can be exported.
	DocumentOutputFormat string
	// ConversionStatus represents the status of a document conversion.
	ConversionStatus struct {
		Status ConversionStatusType `json:"status"`
	}
	// Data represents mathematical expressions in different notation formats.
	// Each Data object contains the expression in a specific format (e.g., ASCII math or LaTeX).
	// Multiple formats can be requested for the same content.
	Data struct {
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
	LineData struct {
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
	WordData struct {
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
	DetectedAlphabet struct {
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
	GeometryData struct {
		// Position contains the pixel coordinates for this geometric element
		Position *Position `json:"position,omitempty"`
		// ShapeList contains all shapes detected in the diagram
		ShapeList []ShapeData `json:"shape_list"`
		// LabelList contains all labels associated with the geometric elements
		LabelList []LabelData `json:"label_list"`
	}
	// Position represents pixel coordinates in the image.
	// The coordinate system starts from the top-left corner of the image.
	Position struct {
		// X coordinate, counting from top left
		X int `json:"x"`
		// Y coordinate, counting from top left
		Y int `json:"y"`
	}
	// ShapeData represents a geometric shape detected in the image.
	// Currently supports triangle detection, with potential for additional
	// shape types in future updates.
	ShapeData struct {
		// Type specifies the type of geometric shape
		// Currently only "triangle" is supported
		Type string `json:"type"`
		// VertexList contains all vertices that make up this shape
		VertexList []VertexData `json:"vertex_list"`
	}
	// VertexData represents a vertex in a geometric shape.
	// It includes both the position of the vertex and its connections
	// to other vertices in the same shape.
	VertexData struct {
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
	LabelData struct {
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
	ErrorInfo struct {
		// Code is a machine-readable error code
		Code string `json:"code"`
		// Message is a human-readable error description
		Message string `json:"message"`
		// Details contains any additional error-specific information
		Details any `json:"details,omitempty"`
	}
	// DataOptions represents configuration for various output formats of image data.
	// All fields are optional and will use their default values if not specified.
	DataOptions struct {
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
	// ConversionFormats specifies which output formats should be generated from the Mathpix Markdown.
	ConversionFormats struct {
		// MMD represents Mathpix Markdown format conversion
		MMD bool `json:"mmd,omitempty"`
		// MD represents standard Markdown format conversion
		MD bool `json:"md,omitempty"`
		// DOCX represents Microsoft Word format conversion
		DOCX bool `json:"docx,omitempty"`
		// TeXZip represents LaTeX with images in ZIP format conversion
		TeXZip bool `json:"tex.zip,omitempty"`
		// HTML represents HTML format conversion
		HTML bool `json:"html,omitempty"`
		// PDFWithHTML represents PDF with HTML rendering conversion
		PDFWithHTML bool `json:"pdf_html,omitempty"`
		// PDFWithLaTeX represents PDF with LaTeX equations conversion
		PDFWithLaTeX bool `json:"pdf_latex,omitempty"`
	}
	// AlphabetsAllowed represents options for specifying which alphabets are allowed in the output.
	AlphabetsAllowed struct {
		// Formats specifies the allowed format types
		Formats []string `json:"formats"`
		// AlphabetsAllowed is a map controlling which alphabets are allowed in the output
		// Keys correspond to alphabet codes (e.g. "hi" for Hindi, "ru" for Russian)
		// false value indicates the alphabet is disabled, true or omission indicates it's allowed
		//
		// By default all alphabets are allowed in the output, to disable alphabet specify "alphabets_allowed": {"alphabet_key": false}.
		// Specifying "alphabets_allowed": {"alphabet_key": true} has the same effect as not specifying that alphabet inside alphabets_allowed map.
		AlphabetsAllowed map[string]bool `json:"alphabets_allowed"`
	}
	// StrokesData contains the actual stroke coordinates
	StrokesData struct {
		// Strokes contains arrays of x and y coordinates representing the handwriting.
		Strokes StrokeCoordinates `json:"strokes"`
	}
	// StrokeCoordinates contains the x and y coordinates for each stroke
	StrokeCoordinates struct {
		// X contains arrays of x-coordinates, where each array represents one stroke.
		X [][]int `json:"x"`
		// Y contains arrays of y-coordinates, where each array represents one stroke.
		Y [][]int `json:"y"`
	}
	// OCRResult represents a single OCR result entry with information about the processing
	OCRResult struct {
		// ISO timestamp of recorded result information
		Timestamp string `json:"timestamp"`

		// API endpoint used for upload (eg `/v3/text`, `/v3/strokes`, ...)
		Endpoint string `json:"endpoint"`

		// Difference between timestamp and when request was received
		Duration float64 `json:"duration"`

		// Request body arguments
		RequestArgs *OCRRequestArgs `json:"request_args"`

		// Result body for request
		Result *ResultBody `json:"result"`

		// An object of detections for each request
		Detections *Detections `json:"detections"`
	}

	// OCRRequestArgs represents the original request arguments
	OCRRequestArgs struct {
		// Tags associated with the request
		Tags []string `json:"tags,omitempty"`

		// Requested output formats
		Formats []string `json:"formats,omitempty"`
	}

	// ResultBody represents the OCR processing result body
	ResultBody struct {
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
	Detections struct {
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
)

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

// String returns the string representation of the ImageFormat
func (f ImageFormat) String() string {
	return string(f)
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

// String returns the string representation of the InputFormat.
// This method satisfies the Stringer interface.
func (f InputFormat) String() string {
	return string(f)
}

// IsValid checks if the input format is supported by comparing against
// known valid formats. Returns true if the format is supported, false otherwise.
func (f InputFormat) IsValid() bool {
	switch f {
	case InputFormatPDF, InputFormatEPUB, InputFormatDOCX, InputFormatPPTX,
		InputFormatAZW, InputFormatAZW3, InputFormatKFX, InputFormatMOBI,
		InputFormatDJVU, InputFormatDOC, InputFormatWPD, InputFormatODT:
		return true
	}
	return false
}

// String returns the string representation of the OutputFormat.
// This method satisfies the Stringer interface.
func (f DocumentOutputFormat) String() string {
	return string(f)
}
