package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
	"os"
)

func main() {

	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	router := NewRouter()

	//LoadConfig does just that, load your config
	LoadConfig()

	//ConnectDB connects to Database
	ConnectDB(DBuser, DBpass, DBhost, DBport, DBname)

	// Setup TLS
	b := []byte(TLSKey)
	b = append([]byte(TLSCert))
	var v *pem.Block
	var pkey []byte
	var pemBlocks []*pem.Block
	for {
		v, b = pem.Decode(b)
		if v == nil {
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			if x509.IsEncryptedPEMBlock(v) {
				pkey, _ = x509.DecryptPEMBlock(v, []byte(TLSPass))
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
			} else {
				pkey = pem.EncodeToMemory(v)
			}
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}
	//Encode Combined and decrypted key to memory
	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
	// Construct a tls.config
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS11,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_AES_256_GCM_SHA384,
		},
		Certificates: []tls.Certificate{c},
	}

	// Build a server:
	server := http.Server{
		// Other options
		TLSConfig:    cfg,
		Handler:      router,
		Addr:         ":9999",
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(server.ListenAndServeTLS("", ""))
}
