package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Listen and serve
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Print the result to the responsewriter

	// check the methode
	switch r.Method {
	case "GET":
		// get the parameter from the get request
		keys := r.URL.Query()

		key := keys["key"]
		fmt.Fprint(w, "Key: ", key[0], "\n")

		api := keys["api"]
		fmt.Fprint(w, "api: ", api[0], "\n")

		fmt.Fprint(w, "Hello World - get")
	case "POST":
		fmt.Fprint(w, "Hello World - post\n")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "Body-Get: %s\n", body)
	default:
		fmt.Fprint(w, "Only GET and PUT are supported")
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
