package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, secure world from server!")
}

func main() {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair("/Users/aadharmishra/Documents/mtls/server.crt", "/Users/aadharmishra/Documents/mtls/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// Create a certificate pool for trusted CA
	caPool := x509.NewCertPool()
	caCert, err := os.ReadFile("/Users/aadharmishra/Documents/mtls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caPool.AppendCertsFromPEM(caCert)

	// Configure TLS server with mutual authentication
	tlsConfig := &tls.Config{
		ClientCAs:    caPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
	}

	// Create listener with TLS config
	listener, err := tls.Listen("tcp", ":8443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create HTTP server with handler and listener
	srv := &http.Server{Handler: http.HandlerFunc(handler)}
	log.Fatal(srv.Serve(listener))
}
