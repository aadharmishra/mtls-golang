package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := []byte(`{"message": "Hello from the server!"}`)
	w.Write(data)
}

func main() {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
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
