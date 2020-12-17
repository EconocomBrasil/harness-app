package main

import (
	"fmt"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Econocom Says Hello!!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	tracer.Start()
	defer tracer.Stop()

	mux := httptrace.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", mux)

}
