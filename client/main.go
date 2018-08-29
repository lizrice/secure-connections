package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lizrice/secure-connections/utils"
)

func main() {
	client := getClient()
	resp, err := client.Get("https://liz-server:8080")
	must(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	must(err)

	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../ca/minica.pem")
	cp.AppendCertsFromPEM(data)

	// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

	config := &tls.Config{
		// Certificates: []tls.Certificate{c},
		RootCAs:               cp,
		GetClientCertificate:  utils.ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}

// fmt.Println("Certificate authority:")
// must(utils.OutputPEMFile("../ca/cert"))
// cp, _ := x509.SystemCertPool() or
// cp := x509.NewCertPool()
// data, _ := ioutil.ReadFile("../ca/cert")
// cp.AppendCertsFromPEM(data)

// fmt.Println("My certificate:")
// must(utils.OutputPEMFile("signed-cert"))
// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

// InsecureSkipVerify: true,
// RootCAs:               cp,
// Certificates:          []tls.Certificate{c},
