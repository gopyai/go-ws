package ws2_test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gopyai/ws2"
)

func Example() {
	mux := http.NewServeMux()
	mux.Handle("/api/", ws2.Handler(f))

	go func() {
		if e := http.ListenAndServe(":8080", mux); e != nil {
			fmt.Println("Error:", e)
			os.Exit(1)
		}
	}()
	time.Sleep(time.Millisecond)

	u, _ := url.Parse("http://localhost:8080/api/hello")
	q := u.Query()
	q.Set("msg", "hello")
	u.RawQuery = q.Encode()

	h, o, s, e := ws2.Call(
		"PUT",
		u.String(),
		"text/plain",
		map[string]string{"Request-Header": "OK"},
		[]byte("Body"),
		3)

	fmt.Println("### RESPONSE ###")
	fmt.Println("Response Header:", h["Response-Header"])
	fmt.Println("Output:", string(o))
	fmt.Println("Status Code:", s)
	fmt.Println("Error:", e)

	// Output:
	// ### REQUEST ###
	// Method: PUT
	// URI: /api/hello?msg=hello
	// Content-Type: text/plain
	// Request Header: map[Request-Header:OK]
	// Input: Body
	// ### RESPONSE ###
	// Response Header: OK
	// Output: OK
	// Status Code: 200
	// Error: <nil>
}

func f(
	method, uri, contentType string, inHeader map[string]string, in []byte) (
	outHeader map[string]string, out []byte, code int, err error) {

	fmt.Println("### REQUEST ###")
	fmt.Println("Method:", method)
	fmt.Println("URI:", uri)
	fmt.Println("Content-Type:", contentType)
	fmt.Println("Request Header:", inHeader)
	fmt.Println("Input:", string(in))

	return map[string]string{"Response-Header": "OK"}, []byte("OK"), 200, nil
}
