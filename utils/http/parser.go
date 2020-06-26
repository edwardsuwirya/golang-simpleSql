package http

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}
func (p *Parser) Form(r *http.Request, i interface{}) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	}
	return nil
}
