package helper

import "strings"

//Response use for json return 
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}


//EmptyObject used when data doesnt want to be nil on json format
type EmptyObject struct{}


//MakeResponse use to make response when API is called
func MakeResponse(status bool, message string, data interface{}) *Response {
	res := &Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return res
}

//MakeErrorResponse used to make error response when calling API
func MakeErrorResponse(message string, err string, data interface{}) *Response {
	error := strings.Split(err, "\n")
	res := &Response{
		Status:  false,
		Message: message,
		Error:   error,
		Data:    data,
	}

	return res

}