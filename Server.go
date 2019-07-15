package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
)

//KeyDecrypt Decrypts an RSA Encrypted KEY and returns a TLS Certificate with unencrypted key and cert pair
func KeyDecrypt() tls.Certificate {
	var pemBlocks []*pem.Block
	certAndKey := TLSKey + TLSCert + TLSChain
	var v *pem.Block
	var pkey []byte
	b := []byte(certAndKey)
	var err error

	for {
		v, b = pem.Decode(b)

		if v == nil {
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			log.Println("Found a private Key")
			if x509.IsEncryptedPEMBlock(v) {
				log.Println("Private Key is encrypted, attempting decryption")
				pkey, err = x509.DecryptPEMBlock(v, []byte(TLSPass))
				if err != nil {
					log.Fatal("Error Decrypting Key")
				}
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
				log.Println("Decrypted Private Key sucessfully")
			} else {
				pkey = pem.EncodeToMemory(v)
			}
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}
	log.Printf("Lenth of pemBlock is %b", len(pemBlocks))
	// chain, rest := pem.Decode([]byte(TLSChain))
	// _ = rest
	// pemBlocks = append(pemBlocks, chain)
	// for i, v := range pemBlocks {
	// 	log.Printf("Pemblock is %v and its element is %s", pemBlocks[i], v)

	// }

	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
	return c
}

//BuildChain Returns a Cert Pool of the CA Certs provided
func BuildChain() x509.CertPool {
	chain := []byte(TLSChain)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(chain)
	return *caCertPool
}

//NewServer returns an HTTP Server Pointer
func NewServer() *http.Server {

	router := NewRouter()
	c := KeyDecrypt()

	//caCerts := BuildChain()
	// chain := []byte(TLSChain)
	// caCertPool := x509.NewCertPool()
	// if caCertPool.AppendCertsFromPEM(chain) {
	// 	log.Print("Loaded CA Cahin")
	// } else {
	// 	log.Println("Could not load CA Chain")
	// }
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM([]byte(TLSChain)); !ok {
		log.Println("Failed to Load CA Certs")
	}

	// Construct a tls.config
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			// tls.TLS_AES_128_GCM_SHA256,
			// tls.TLS_AES_256_GCM_SHA384,
			// tls.TLS_CHACHA20_POLY1305_SHA256,
			// tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			// tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			// tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			// tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			// tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			// tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			// tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{c},
	}
	// Build a server:
	server := http.Server{
		// Other options
		TLSConfig:    cfg,
		Handler:      router,
		Addr:         ":" + ServerPort,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Println("TLS Enabled WebServer Started")
	return &server
}
