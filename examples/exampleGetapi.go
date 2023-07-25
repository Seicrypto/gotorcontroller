package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/Seicrypto/gotorcontroller"
	// "../" should be:
	// "github.com/Seicrypto/gotorcontroller"
	// in your project
)

func TorGet(funcURL string) string {
	// Set proxy, this sample was working in docker golang:bullseye
	// Using 8118 port 'cause torcontroller setting provixy config to 8118 port.
	proxyURL, err := url.Parse("http://127.0.0.1:8118")
	if err != nil {
		fmt.Println("Failed to parse proxy URL:", err)
		os.Exit(1)
	}
	// build HTTP client
	httpTorClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	// Send GET require.
	resp, err := httpTorClient.Get(funcURL)
	if err != nil {
		fmt.Println("Failed to send GET request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		os.Exit(1)
	}
	// Print the response body.
	return string(body)
}

func main() {
	// Start tor service.
	gotorcontroller.StartTor()
	fmt.Println(TorGet("https://api64.ipify.org?format=json"))
	//{"ip":"23.137.249.150"}
	gotorcontroller.SwitchIP()
	fmt.Println(TorGet("https://api64.ipify.org?format=json"))
	//{"ip":"2.58.56.37"}
}
