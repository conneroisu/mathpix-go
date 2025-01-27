package mathpix

import "time"

// https://docs.mathpix.com/?shell#query-ocr-usage

type usagePayload struct {
	Payload *RequestUsage `json:"payload"`
}

// RequestUsage is the payload for the request to get the ocr usage of the API.
type RequestUsage struct {
	FromDate time.Time `in:"query=from_date"`
	ToDate   time.Time `in:"query=to_date"`
	GroupBy  string    `in:"query=group_by,required"`
	Timespan string    `in:"query=timespan,required"`
}

// ResponseUsage is the response for the request to get the ocr usage of the API.
type ResponseUsage struct {
	OcrUsage []struct {
		FromDate        time.Time `json:"from_date"`
		AppID           []string  `json:"app_id"`
		UsageType       string    `json:"usage_type"`
		RequestArgsHash []string  `json:"request_args_hash"`
		Count           int       `json:"count"`
	} `json:"ocr_usage"`
}
