package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{}

	// Create a pool with the certificate, since it is not signed by any CA
	caCert, err := ioutil.ReadFile("../../cert/server.crt")
	if err != nil {
		log.Fatalf("Error reading certificate. %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with certificate from server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Perform the request
	resp, err := client.Post("https://localhost:9191/hello/sayHello", "text/plain", bytes.NewBufferString("Hello Go!"))
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))
}
