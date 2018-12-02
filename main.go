package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Configuration struct {
	Listen string `json:"listen"`
	Whitelist []string `json:"whitelist"`
}

var configuration = Configuration{":9999", []string{"^https:\\/\\/dynamic\\.lunanode\\.com"}}

const configFile = "tsconf.json"

func readConfiguration() {

	file, err := os.Open(configFile)
	defer file.Close()

	if err != nil {
		fmt.Println("No tsconf.json found. ")
		return
	}

	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println( "Whitelist:" )

	for _, pattern := range configuration.Whitelist {
		fmt.Println( "  * "+pattern )
	}

}

func forward(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Entering 'forward'")

	contentLength, err := strconv.ParseInt(req.Header.Get("content-length"),10,64)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(503)
		fmt.Println("...done (503)")
		return
	}

	url := req.URL.RawQuery

	processRequest := false

	for _, pattern := range configuration.Whitelist {

		match, err := regexp.MatchString(pattern, url)

		if err != nil {
			// ignore pattern
			continue
		}

		if match {
			processRequest = true
			break
		}

	}

	if !processRequest {
		fmt.Printf("%s not in whitelist. Rejecting.\n", url)
		w.WriteHeader(404)
		fmt.Println("...done (404)")
		return
	}

	origin := req.Header.Get("origin")

	req.Header.Del("origin")
	req.Header.Del("host")
	req.Header.Del("content-length")



	proxyReq, err := http.NewRequest("POST", url, req.Body)

	if err != nil {
		fmt.Println(err);
		w.WriteHeader(503)
		fmt.Println("...done (503)")
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
		fmt.Println("...done (503)")
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
		fmt.Println("...done (503)")
		return
	}

	fmt.Println("[RX]" )
	w.Write(bodyBytes)

	fmt.Println("...done (200)")

}
func main() {
	readConfiguration()
	http.HandleFunc("/", forward)
	fmt.Printf("Starting proxy on %s\n", configuration.Listen)
	if err := http.ListenAndServe(configuration.Listen, nil); err != nil {
		panic(err)
	}
}