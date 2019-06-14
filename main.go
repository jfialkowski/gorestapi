package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte
	var b []byte
	fmt.Printf("TLSKey is %v \n", TLSKey)
	//for {
	v, b = pem.Decode(TLSKey)
	fmt.Printf("V is %v \n", v)
	if v == nil {
		log.Fatal("V is Nil")
		//break
	}
	if v.Type == "PRIVATE KEY" {
		fmt.Println("found private key")
		if x509.IsEncryptedPEMBlock(v) {
			pkey, _ = x509.DecryptPEMBlock(v, []byte(TLSPass))
			pkey = pem.EncodeToMemory(&pem.Block{
				Type:  v.Type,
				Bytes: pkey,
			})
		} else {
			fmt.Println("Encoded to memory")
			pkey = pem.EncodeToMemory(v)
		}
	} else {
		fmt.Printf("found %v in rest \n ", b)
		pemBlocks = append(pemBlocks, v)
	}
	//}
	//Encode Combined and decrypted key to memory

	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), b)

	// Construct a tls.config
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
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
