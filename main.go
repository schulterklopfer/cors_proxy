package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func forward(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Entering 'forward'")

	contentLength, err := strconv.ParseInt(req.Header.Get("content-length"),10,64)

	if err != nil {
		fmt.Println(err);
		w.WriteHeader(503)
		return
	}

	origin := req.Header.Get("origin")

	req.Header.Del("origin")
	req.Header.Del("host")
	req.Header.Del("content-length")


	url := req.URL.RawQuery

	proxyReq, err := http.NewRequest("POST", url, req.Body)

	if err != nil {
		fmt.Println(err);
		w.WriteHeader(503)
		return
	}

	proxyReq.ContentLength = contentLength;

	for header, _ := range req.Header {
		proxyReq.Header.Set(header,req.Header.Get(header))
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)

	if err != nil {
		fmt.Println(err);
		w.WriteHeader(503)
		return
	}

	fmt.Printf("[TX] %s\n", url)

	for header, _ := range resp.Header {
		w.Header().Set(header,resp.Header.Get(header))
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true" )
	w.Header().Set("Access-Control-Allow-Origin", origin)

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err);
		w.WriteHeader(503)
		return
	}

	fmt.Println("[RX]" )
	w.Write(bodyBytes)

	fmt.Println("...done")

}
func main() {
	http.HandleFunc("/", forward)
	fmt.Println("Starting proxy on port 9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}