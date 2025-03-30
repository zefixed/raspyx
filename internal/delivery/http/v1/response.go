package v1

type ResponseOK struct {
	Status   string      `json:"status" example:"OK"`
	Response interface{} `json:"response,omitempty"`
}

type ResponseError struct {
	Status string `json:"status"  example:"Error"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func RespOK(response interface{}) ResponseOK {
	return ResponseOK{
		Status:   StatusOK,
		Response: response,
	}
}

func RespError(msg string) ResponseError {
	return ResponseError{
		Status: StatusError,
		Error:  msg,
	}
}
