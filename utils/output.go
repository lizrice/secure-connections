package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// CertificateInfo returns a string describing the certificate
func CertificateInfo(cert *x509.Certificate) string {
	if cert.Subject.CommonName == cert.Issuer.CommonName {
		return fmt.Sprintf("    Self-signed certificate %v\n", cert.Issuer.CommonName)
	}

	s := fmt.Sprintf("    Subject %v\n", cert.DNSNames)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.CommonName)
	return s
}

// CertificateChains prints information about verified certificate chains
func CertificateChains(rawCerts [][]byte, chains [][]*x509.Certificate) error {
	if len(chains) > 0 {
		fmt.Println("Verified certificate chain from peer:")

		for _, v := range chains {
			// fmt.Printf("Chain %d:\n", j)
			for i, cert := range v {
				fmt.Printf("  Cert %d:\n", i)
				fmt.Printf(CertificateInfo(cert))
			}
		}
		Wait()
	}

	return nil
}

// OutputPEMFile reads info from a PEM file and displays it
func OutputPEMFile(filename string) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	for len(data) > 0 {
		var block *pem.Block
		block, data = pem.Decode(data)
		fmt.Printf("Type: %#v\n", block.Type)
		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return err
			}
			fmt.Printf(CertificateInfo(cert))
		default:
			fmt.Println(block.Type)
		}
	}
	return nil
}
