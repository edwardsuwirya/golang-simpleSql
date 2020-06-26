package http

import (
	"github.com/edwardsuwirya/simpleSql/utils/message"
	"net/http"
)

type IResponder interface {
	Data(w http.ResponseWriter, status int, message string, data interface{})
	Error(w http.ResponseWriter, status int, error message.IErrorContent)
}

type ErrorResponse struct {
	ErrorID int    `json:"errorId"`
	Message string `json:"message"`
}

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
