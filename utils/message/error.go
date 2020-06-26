package message

import "github.com/edwardsuwirya/simpleSql/utils/appStatus"

type IErrorContent interface {
	Code() int
	Message() string
}

type ErrorMessage struct {
	ErrorID int    `json:"errorId"`
	Msg     string `json:"message"`
}

func NewErrorMessage(code int) *ErrorMessage {
	return &ErrorMessage{
		ErrorID: code,
		Msg:     appStatus.StatusText(code),
	}
}

func (c *ErrorMessage) Code() int {
	return c.ErrorID
}

func (c *ErrorMessage) Message() string {
	return c.Msg
}
