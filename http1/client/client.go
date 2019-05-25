package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed by a known CA
	caCert, err := ioutil.ReadFile("../../cert/server.crt")

	if err != nil {
		log.Fatalf("Error reading server certificate. %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Perform the request
	res, err := client.Post("https://localhost:9191/hello/sayhello", "text/plain", bytes.NewBufferString("Hello Go!"))
	if err != nil {
		log.Fatalf("Request failed. %s", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed reading response body. %s", err)
	}

	fmt.Printf("Status Code: %d, Proto: %s, Body: %s", res.StatusCode, res.Proto, string(body))
}
