package mathpix

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/ggicci/httpin"
)

// Client is the main struct for the mathpix-go library.
type (
	Client struct {
		apiKey  string
		appID   string
		baseURL url.URL
		client  *http.Client
		logger  *slog.Logger

		SetCommonHeaders func(req *http.Request)
	}

	// ClientOption is a function that can be used to configure a Client.
	ClientOption func(*Client)
)

// NewClient creates a new Client with the given API key and base URL.
func NewClient(apiKey, appID string, opts ...ClientOption) *Client {
	baseURL, _ := url.Parse("https://api.mathpix.com")
	client := &Client{
		apiKey:  apiKey,
		appID:   appID,
		baseURL: *baseURL,
		client:  http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}
	client.SetCommonHeaders = func(req *http.Request) {
		req.Header.Set("app_key", client.apiKey)
		req.Header.Set("app_id", client.appID)
	}
	return client
}

// WithBaseURL sets the base URL for the Client.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if u, err := url.Parse(baseURL); err == nil {
			c.baseURL = *u
		}
	}
}

// WithLogger sets the logger for the Client.
func WithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) { c.logger = logger }
}

// WithClient sets the client for the Client.
func WithClient(client *http.Client) ClientOption {
	return func(c *Client) { c.client = client }
}

// call is a method that takes an endpoint and make a call to it.
func call[Request, Response any](
	ctx context.Context,
	c *Client,
	e endpoint[Request, Response],
	request Request,
	param string,
) (response Response, err error) {
	httpReq, err := httpin.NewRequestWithContext(
		ctx,
		e.method,
		c.baseURL.JoinPath(e.name, param).String(),
		request,
	)
	if err != nil {
		return response, err
	}
	c.SetCommonHeaders(httpReq)
	contentType := httpReq.Header.Get("Content-Type")
	if contentType == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	res, err := c.client.Do(httpReq)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resp, apiErr, err := nopDecode[APIError](res.Body)
	if err != nil {
		return response, err
	}
	if res.StatusCode < http.StatusOK ||
		res.StatusCode >= http.StatusBadRequest ||
		isErrorID(apiErr.ID) {
		return response, &apiErr
	}
	err = json.NewDecoder(resp).Decode(&response)
	if err != nil {
		return
	}
	return response, nil
}

// nopDecode decodes the request body into the given type.
func nopDecode[T any](
	r io.Reader,
) (io.Reader, T, error) {
	var (
		buf bytes.Buffer
		v   T
	)
	tee := io.TeeReader(r, &buf)
	err := json.NewDecoder(tee).Decode(&v)
	if err != nil {
		return nil, v, err
	}
	return &buf, v, nil
}

// Image sends an image to the Mathpix API.
func (c *Client) Image(
	ctx context.Context,
	request *ImageRequest,
) (*ImageResponse, error) {
	return call(
		ctx,
		c,
		imagesEndpoint,
		&imageRequestPayload{
			Payload: request,
		},
		"",
	)
}

// Pdf sends a PDF to the Mathpix API.
func (c *Client) Pdf(
	ctx context.Context,
	request *RequestDocument,
) (*DocumentResponse, error) {
	return call(
		ctx,
		c,
		documentsEndpoint,
		&documentRequestPayload{
			Payload: request,
		},
		"",
	)
}

// PdfResult represents the result of a PDF Result request.
func (c *Client) PdfResult(
	ctx context.Context,
	request *ResultRequest,
) (*ConversionResultResponse, error) {
	return call(
		ctx,
		c,
		conversionStatusEndpoint,
		nil,
		request.PDFID,
	)
}

// Batch sends a batch of images to the Mathpix API.
func (c *Client) Batch(
	ctx context.Context,
	request *RequestPostBatch,
) (*PostBatchResponse, error) {
	return call(
		ctx,
		c,
		batchEndpoint,
		&postBatchRequestPayload{
			Payload: request,
		},
		"",
	)
}

// RequestStrokes sends a strokes recognition request to the Mathpix API.
func (c *Client) RequestStrokes(
	ctx context.Context,
	request *RequestStrokes,
) (*StrokesResponse, error) {
	return call(
		ctx,
		c,
		strokesEndpoint,
		&requestStrokesPayload{
			Payload: request,
		},
		"",
	)
}

// SearchResults searches for OCR results.
func (c *Client) SearchResults(
	ctx context.Context,
	request *OCRSearchRequest,
) (*OCRResultsResponse, error) {
	return call(
		ctx,
		c,
		ocrResultsEndpoint,
		&ocrResultsPayload{
			Payload: request,
		},
		"",
	)
}

// NewClientToken creates a new temporary app token.
func (c *Client) NewClientToken(
	ctx context.Context,
	request *AppTokenRequest,
) (*AppTokenResponse, error) {
	return call(
		ctx,
		c,
		appTokensEndpoint,
		&appTokenPayload{
			Payload: request,
		},
		"",
	)
}

// GetBatch retrieves a batch of images.
func (c *Client) GetBatch(
	ctx context.Context,
	batchID string,
) (*GetBatchResponse, error) {
	return call(
		ctx,
		c,
		getBatchEndpoint,
		&getBatchPayload{},
		batchID,
	)
}

// RequestUsage sends a request to get the ocr usage of the API.
func (c *Client) RequestUsage(
	ctx context.Context,
	request *RequestUsage,
) (*UsageResponse, error) {
	return call(
		ctx,
		c,
		requestUsageEndpoint,
		&usagePayload{
			Payload: request,
		},
		"",
	)
}
