package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var request string
	url := fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto)
	request += url
	request += fmt.Sprintf("Host: %v\n", r.Host)
	for name, headers := range r.Header {
		// name = strings.ToLower(name)
		for _, h := range headers {
			request += fmt.Sprintf("%v: %v\n", name, h)
		}
	}

	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	request += fmt.Sprintf("Body: %v\n", string(bodyBuffer))

	fmt.Fprintf(w, "%s", request)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
