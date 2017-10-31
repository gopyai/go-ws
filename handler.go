package ws2

import (
	"io/ioutil"
	"net/http"
)

type (
	WsFunc func(
		method, uri, contentType string, inHeader map[string]string, in []byte) (
		outHeader map[string]string, out []byte, code int, err error)
)

func Handler(f WsFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var e error
		var c int
		var in []byte
		if in, e = ioutil.ReadAll(r.Body); e == nil {
			var h map[string]string
			var o []byte
			if h, o, c, e = f(
				r.Method,
				r.RequestURI,
				r.Header.Get("Content-Type"),
				convertHttpHeader(r.Header),
				in); e == nil {
				for k, v := range h {
					w.Header().Set(k, v)
				}
				w.WriteHeader(c)
				if _, e = w.Write(o); e == nil {
					return
				}
			}
		}
		if c == 0 || c == http.StatusOK {
			c = http.StatusInternalServerError
		}
		http.Error(w, e.Error(), c)

	})
}
