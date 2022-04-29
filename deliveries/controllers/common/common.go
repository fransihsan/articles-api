package common

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InternalServerError(message string) Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}


