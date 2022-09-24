package utils

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

func NewResponse(statusCode int, data interface{}, message interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
