package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func main() {

	caCert, err := ioutil.ReadFile("../certs/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPoll := x509.NewCertPool()
	caCertPoll.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPoll,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(server.ListenAndServeTLS("../certs/cert.pem", "../certs/key.pem"))
}
