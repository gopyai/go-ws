package ws2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Call URL. The timeout is in seconds.
// There are 2 kinds of error:
// - ErrCall is for general error.
// - ErrStatus is when the HTTP status is not 200 (OK).
func Call(
	method, url, contentType string,
	inHeader map[string]string,
	inB []byte,
	timeout int,
) (
	outHeader map[string]string,
	outB []byte,
	statusCode int,
	err error,
) {

	// Create http client and request
	cli := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	req, e := http.NewRequest(method, url, bytes.NewReader(inB))
	if e != nil {
		return nil, nil, 0, ErrCall
	}

	// Set proper header
	if inHeader != nil {
		for k, v := range inHeader {
			req.Header[k] = []string{v}
		}
	}
	req.Header["Content-Type"] = []string{contentType} // Overwrite

	// Make request
	res, e := cli.Do(req)
	if e != nil {
		return nil, nil, 0, ErrCall
	}
	defer res.Body.Close()

	// Convert response header
	resHeader := convertHttpHeader(res.Header)

	// Read body
	out, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, nil, 0, ErrCall
	}

	return resHeader, out, res.StatusCode, e
}
