package helper

import "strings"

// Response is used for static shape json return
type ResponseLogin struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
	Data    interface{} `json:"data"`
	Token   interface{} `json:"token"`
}
type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

// Response Success
func BuildResponse(status int64, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// Response Error
func BuildErrorResponse(status int64, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "/n")
	res := Response{
		Status:  status,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func BuildValidationResponse(e interface{}) interface{} {
	res := map[string]interface{}{"validationError": e}
	return res
}
