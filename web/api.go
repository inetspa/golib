package web

type ApiResult struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Total   int         `json:"total,omitempty"`
	Count   int         `json:"count,omitempty"`
}
