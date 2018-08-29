package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/lizrice/secure-connections/utils"
)

func main() {
	server := getServer()
	http.HandleFunc("/", myHandler)
	must(server.ListenAndServe())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")
	w.Write([]byte("Hey GopherCon!"))
}

func getServer() *http.Server {

	tls := &tls.Config{
		GetCertificate:        utils.CertReqFunc("", ""),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tls,
	}
	return server
}

func must(err error) {
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}

// NoClientCert ClientAuthType = iota
// RequestClientCert
// RequireAnyClientCert
// VerifyClientCertIfGiven
// RequireAndVerifyClientCert
