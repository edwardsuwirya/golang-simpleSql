package http

import (
	"encoding/json"
	"github.com/edwardsuwirya/simpleSql/utils/message"
	"net/http"
	"strconv"
)

type jsonResponder struct {
}

func NewDefaultJSONResponder() IResponder {
	return &jsonResponder{}
}

func (j *jsonResponder) Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}

	content, _ := json.Marshal(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	_, _ = w.Write(content)
}
func (j *jsonResponder) Data(w http.ResponseWriter, status int, message string, data interface{}) {
	content := DataResponse{Message: message, Data: data}
	j.Write(w, status, content)

}
func (j *jsonResponder) Error(w http.ResponseWriter, status int, error message.IErrorContent) {
	content := ErrorResponse{ErrorID: error.Code(), Message: error.Message()}
	j.Write(w, status, content)
}
