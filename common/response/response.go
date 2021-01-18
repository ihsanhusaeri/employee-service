package response

//Response defines the response format
type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Payload    interface{} `json:"payload"`
}

func NewResponse(statusCode int, message interface{}, payload interface{}) (int, *Response) {
	err, ok := message.(error)

	if ok {
		message = err.Error()
	}

	return statusCode, &Response{
		StatusCode: statusCode,
		Message:    message,
		Payload:    payload,
	}
}
