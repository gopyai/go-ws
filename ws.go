package ws2

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrCall = errors.New("Web Service call error")
)

func convertHttpHeader(h http.Header) map[string]string {
	m := make(map[string]string)
	for k, v := range h {
		switch k {
		case "Accept":
		case "Accept-Encoding":
		case "Accept-Language":
		case "Cache-Control":
		case "Connection":
		case "Content-Length":
		case "Content-Type":
		case "Origin":
		case "User-Agent":
		default:
			m[k] = strings.Join(v, ",")
		}
	}
	return m
}
