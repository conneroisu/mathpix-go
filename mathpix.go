package mathpix

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"path"

	"github.com/ggicci/httpin"
)

// Client is the main struct for the mathpix-go library.
type (
	Client struct {
		appKey  string
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
func NewClient(apiKey string, opts ...ClientOption) *Client {
	baseURL, _ := url.Parse("https://api.mathpix.com")
	client := &Client{
		appKey:  apiKey,
		baseURL: *baseURL,
	}
	for _, opt := range opts {
		opt(client)
	}
	client.SetCommonHeaders = func(req *http.Request) {
		req.Header.Set("app_key", client.appKey)
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
) (*ResponseDocument, error) {
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

// Batch sends a batch of images to the Mathpix API.
func (c *Client) Batch(
	ctx context.Context,
	request *RequestPostBatch,
) (*ResponsePostBatch, error) {
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
	return call(ctx, c, strokesEndpoint, &requestStrokesPayload{
		Payload: request,
	}, "")
}

// SearchResults searches for OCR results.
func (c *Client) SearchResults(
	ctx context.Context,
	request *SearchParams,
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
) (*ResponseUsage, error) {
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
		path.Join(c.baseURL.Path, e.name, param),
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
	if res.StatusCode < http.StatusOK ||
		res.StatusCode >= http.StatusBadRequest {
		return response, handleErrorResp(res, request, response)
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return
	}
	return response, nil
}

func handleErrorResp[
	Request,
	Response any,
](
	_ *http.Response,
	_ Request,
	_ Response,
) error {
	return nil
}
