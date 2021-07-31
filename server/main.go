package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const contentType = "application/json"

func newServer(port string) *http.Server {

	// Read the Cert Authority file
	caCert, err := ioutil.ReadFile("../certs/demoCA.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Tell the server to use this authority to validate requests instead of the OS default
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":" + port,
		TLSConfig:    tlsConfig,
	}
	return server
}

func fooHandler(w http.ResponseWriter, r *http.Request) {

	// Send a simple json payload
	w.Header().Set("Content-Type", contentType)
	io.WriteString(w, `{"msg":"you have reached foo server"}`)
}

func main() {

	// Setup a server to listen on port 8080
	server := newServer("8080")

	// Set up the server to handle requests
	http.HandleFunc("/foo", fooHandler)
	log.Fatal(server.ListenAndServeTLS("../certs/localhost.crt", "../certs/localhost.key"))
}
