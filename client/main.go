package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := getClient()
	resp, _ := client.Get("http://liz-server:8080")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
}

func getClient() *http.Client {

	config := &tls.Config{
		// GetClientCertificate:  utils.ClientCertReqFunc("",""),
		// VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

// cp, _ := x509.SystemCertPool() or
// cp := x509.NewCertPool()
// data, _ := ioutil.ReadFile("../ca/minica.pem")
// cp.AppendCertsFromPEM(data)

// fmt.Println("My certificate:")
// must(utils.OutputPEMFile("signed-cert"))
// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

// InsecureSkipVerify: true,
// RootCAs:               cp,
// Certificates:          []tls.Certificate{c},

// func must(err error) {
// 	if err != nil {
// 		fmt.Printf("Client error: %v\n", err)
// 		os.Exit(1)
// 	}
// }
