package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
)

//KeyDecrypt Decrypts an RSA Encrypted KEY and returns a TLS Certificate with unencrypted key and cert pair
//func KeyDecrypt() tls.Certificate {
// var pemBlocks []*pem.Block
// certAndKey := TLSKey + "\n" + TLSCert
// log.Printf("KEY is: %s", TLSKey)
// log.Printf("CERT is: %s", TLSCert)
// var v *pem.Block
// var pkey []byte
// b := []byte(certAndKey)
// var err error

// for {
// 	v, b = pem.Decode(b)
// 	if v == nil {
// 		log.Fatal("Could not load a private Key")
// 	}
// 	if v.Type == "RSA PRIVATE KEY" {
// 		log.Println("Found a private Key")
// 		if x509.IsEncryptedPEMBlock(v) {
// 			log.Println("Private Key is encrypted, attempting decryption")
// 			pkey, err = x509.DecryptPEMBlock(v, []byte(TLSPass))
// 			if err != nil {
// 				log.Fatal("Error Decrypting Key")
// 			}
// 			pkey = pem.EncodeToMemory(&pem.Block{
// 				Type:  v.Type,
// 				Bytes: pkey,
// 			})
// 			log.Println("Decrypted Private Key encoded to memory")
// 		} else {
// 			pkey = pem.EncodeToMemory(v)
// 			log.Println("Private Key encoded to memory")
// 		}
// 	} else {
// 		pemBlocks = append(pemBlocks, v)
// 		log.Println("Private Key Simply Appended to PemBlocks")
// 	}
// }
// //log.Printf("pemBlock[0] is: %+v", pemBlocks[0])
// //log.Printf("pkey is is: %+v", pkey)
// c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
//return c
//}

//NewServer returns an HTTP Server Pointer
func NewServer() *http.Server {

	router := NewRouter()
	//c := KeyDecrypt()
	var pemBlocks []*pem.Block
	certAndKey := TLSKey + "\n" + TLSCert
	log.Printf("KEY is: %s", TLSKey)
	log.Printf("CERT is: %s", TLSCert)
	var v *pem.Block
	var pkey []byte
	b := []byte(certAndKey)
	var derr error

	for {
		v, b = pem.Decode(b)
		if v == nil {
			log.Println("Could not load a private Key")
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			log.Println("Found a private Key")
			if x509.IsEncryptedPEMBlock(v) {
				log.Println("Private Key is encrypted, attempting decryption")
				pkey, derr = x509.DecryptPEMBlock(v, []byte(TLSPass))
				if derr != nil {
					log.Println("Error Decrypting Key")
				}
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
				log.Println("Decrypted Private Key encoded to memory")
			} else {
				pkey = pem.EncodeToMemory(v)
				log.Println("Private Key encoded to memory")
			}
		} else {
			pemBlocks = append(pemBlocks, v)
			log.Println("Private Key Simply Appended to PemBlocks")
		}
	}
	//log.Printf("pemBlock[0] is: %+v", pemBlocks[0])
	//log.Printf("pkey is is: %+v", pkey)
	c, err := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
	if err != nil {
		log.Fatal(err)
	}
	//Load Up The TLS Key and Cert to be used in HTTP Daemon
	// var pemBlocks []*pem.Block
	// var v *pem.Block
	// var cert *pem.Block
	// var pkey []byte

	// cert, rest := pem.Decode(TLSCert)
	// if cert == nil {
	// 	log.Fatal("Could not load certificate ")
	// 	fmt.Printf("%v", rest)
	// } else {
	// 	pemBlocks = append(pemBlocks, cert)
	// }
	// v, b := pem.Decode(TLSKey)
	// if v == nil {
	// 	log.Fatal("Could not load Private Key")
	// 	fmt.Printf("%v", b)

	// }
	// if v.Type == "PRIVATE KEY" {
	// 	if x509.IsEncryptedPEMBlock(v) {
	// 		pkey, _ = x509.DecryptPEMBlock(v, []byte(TLSPass))
	// 		pkey = pem.EncodeToMemory(&pem.Block{
	// 			Type:  v.Type,
	// 			Bytes: pkey,
	// 		})
	// 	} else {
	// 		//fmt.Println("Encoded to memory")
	// 		pkey = pem.EncodeToMemory(v)
	// 	}
	// } else {
	// 	pemBlocks = append(pemBlocks, v)
	// }

	// //Encode Combined and decrypted key to memory
	// c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), TLSKey)

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
		Addr:         ":" + ServerPort,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Println("TLS Enabled WebServer Started")
	return &server
}
