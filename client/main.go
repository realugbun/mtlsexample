package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

// Variables holding the TLS cert and Cert Authority
var (
	caCertPool *x509.CertPool
	cert       tls.Certificate
)

// ResponseRec used for unmarshaling json response
type ResponseRec struct {
	Message string `json:"msg"`
}

// FooClient is an example client
type FooClient struct {
	client *http.Client
}

// NewFooClient creates a new instance of the client with the required TLS info
func NewFooClient() *FooClient {

	var newClient FooClient

	newClient.client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	return &newClient
}

func init() {
	var err error

	// Load the cert and key into memory
	cert, err = tls.LoadX509KeyPair("../certs/localhost.crt", "../certs/localhost.key")
	if err != nil {
		log.Fatal(err)
	}

	// Setup our own certificate authority and tell the client to use it instead of the OS default
	caCert, err := ioutil.ReadFile("../certs/demoCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool = x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
}

func main() {

	// Create a new client with all the TLS info attached
	foo := NewFooClient()

	// Send our request over TLS
	r, err := foo.client.Get("https://localhost:8080/foo")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response
	body := ResponseRec{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	// Do somthing
	fmt.Println(body.Message)
}
