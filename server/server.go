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
	caCert, err := ioutil.ReadFile("../certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPoll := x509.NewCertPool()
	caCertPoll.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs: caCertPoll,
		// ClientAuth: tls.RequireAndVerifyClientCert,
		ClientAuth: tls.RequireAndVerifyClientCert,
		ServerName: "localhost",
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// log.Fatal(http.ListenAndServeTLS(":8443", "../certs/server.crt", "../certs/server.key", nil))
	log.Fatal(server.ListenAndServeTLS("../certs/server.crt", "../certs/server.key"))
}
