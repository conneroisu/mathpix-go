package mathpix

type postBatchRequestPayload struct {
	Payload *RequestPostBatch `in:"body=json"`
}

// RequestPostBatch is the request body for the POST /v3/batch endpoint.
//
// The request body may contain any /v3/latex parameters except src and must also contain a urls parameter.
//
// The request may also contain an additional callback parameter to receive results after all the images in the batch have been processed.
type RequestPostBatch struct {
	URLs map[string]string `json:"urls"`
	OCR  string            `json:"ocr_behavior,omitempty"`
	// Callback *Callback         `json:"callback,omitempty"`
}

// ResponsePostBatch is the response from the batch endpoint.
//
// The response contains only a unique batch_id value.
// Even if the request includes a callback, there is no guarantee the callback will run successfully (because of a transient network failure, for example).
// The preferred approach is to wait an appropriate length of time (about one second for every five images in the batch) and then do a GET on /v3/batch/:id where :id is the batch_id value.
// The GET request must contain the same app_id and app_key headers as the POST to /v3/batch.
type ResponsePostBatch struct {
	BatchID string `json:"batch_id"`
}

type getBatchPayload struct{}

// GetBatchResponse is the response from the GET /v3/batch/:id endpoint.
type GetBatchResponse struct {
	Keys    []string               `json:"keys"`
	Results map[string]interface{} `json:"results"`
}
