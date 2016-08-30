package tls

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
)

// NewCACert creates a new Certificate Authority and returns it's private key and public certificate.
func NewCACert(csrFile string) (key, cert []byte, err error) {
	// Open CSR file
	f, err := os.Open(csrFile)
	if os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("%q does not exist", csrFile)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("error opening %q", csrFile)
	}
	// Create CSR struct
	csr := &csr.CertificateRequest{
		KeyRequest: csr.NewBasicKeyRequest(),
	}
	err = json.NewDecoder(f).Decode(csr)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding CSR: %v", err)
	}
	// Generate CA Cert according to CSR
	cert, _, key, err = initca.New(csr)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating CA cert: %v", err)
	}

	return key, cert, nil
}