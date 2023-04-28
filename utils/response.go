package utils

type Response struct {
	Status     string `json:"status"`
	Result     any    `json:"result"`
	Meta       any    `json:"meta,omitempty"`
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code,omitempty"`
}
