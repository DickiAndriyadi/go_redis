package constant

type (
	// Response is used a structure of response
	Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}
)
