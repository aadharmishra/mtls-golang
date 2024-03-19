package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	// Create a certificate pool for trusted CA
	caPool := x509.NewCertPool()
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caPool.AppendCertsFromPEM(caCert)

	// Configure TLS client config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Set to true for development only, skip verification!
		RootCAs:            caPool,
		Certificates:       []tls.Certificate{cert},
	}

	// Create HTTP client with TLS config
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}

	// Make a request to the server
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
