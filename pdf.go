package mathpix

type documentRequestPayload struct {
	Payload *RequestDocument `in:"body=json"`
}

// RequestDocument represents the request parameters for processing a PDF file or URL.
type RequestDocument struct {
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

// ResponseDocument represents the response from the PDF processing endpoint.
type ResponseDocument struct {
	// PDFID is the tracking ID to get status and result
	PDFID string `json:"pdf_id"`
	// Error contains US locale error message if present
	Error string `json:"error,omitempty"`
	// ErrorInfo contains detailed error information
	ErrorInfo map[string]interface{} `json:"error_info,omitempty"`
}

// InputFormat represents supported input file formats for Mathpix processing.
// It is used to specify the format of documents being submitted for processing.
type InputFormat string

// Input format constants define all supported document input types.
const (
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
)

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

// DocumentOutputFormat represents supported output file formats from Mathpix processing.
// It defines the format in which processed documents can be exported.
type DocumentOutputFormat string

// Output format constants define all supported document output types.
const (
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

// String returns the string representation of the OutputFormat.
// This method satisfies the Stringer interface.
func (f DocumentOutputFormat) String() string {
	return string(f)
}

// IsValid checks if the output format is supported by comparing against
// known valid formats. Returns true if the format is supported, false otherwise.
func (f DocumentOutputFormat) IsValid() bool {
	switch f {
	case DocumentFormatMMD, DocumentFormatMD, DocumentFormatDOCX,
		DocumentFormatLaTeXZip, DocumentFormatHTML,
		DocumentFormatPDFWithHTML, DocumentFormatPDFWithLaTeX:
		return true
	}
	return false
}

// AlphabetsAllowed represents options for specifying which alphabets are allowed in the output.
type AlphabetsAllowed struct {
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

// ConversionFormats specifies which output formats should be generated from the Mathpix Markdown.
type ConversionFormats struct {
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
