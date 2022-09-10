package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

const environment = "development"

func main() {
	h := makeHTTPServer()

	if environment == "production" {
		m := &autocert.Manager{
			Cache:      autocert.DirCache("secret-dir"),
			Prompt:     autocert.AcceptTOS,
			Email:      "jason.onggo@tempatkerja.com",
			HostPolicy: autocert.HostWhitelist("estatement.godata.id"),
		}

		hs := makeHTTPServer()
		hs.Addr = ":https"
		hs.TLSConfig = m.TLSConfig()

		go func() {
			log.Printf("Starting HTTPS Server on port %s\n", hs.Addr)
			err := hs.ListenAndServeTLS("", "")
			if err != nil {
				log.Panic(err)
			}
		}()

		h = makeHTTPServerRedirectToHTTPS()

		// allow autocert handle Let's Encrypt auth callbacks over HTTP.
		// it'll pass all other urls to our handler
		h.Handler = m.HTTPHandler(h.Handler)
	} else if environment == "development" {
		h.Addr = ":8080"
	}

	log.Printf("Starting HTTP Server on port %s", h.Addr)
	err := h.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	return makeHTTPServerWithMux(mux)
}

func makeHTTPServerWithMux(mux *http.ServeMux) *http.Server {
	// set timeouts so that a slow or malicious client
	// doesn't hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
		Addr:         ":http",
	}
}

func makeHTTPServerRedirectToHTTPS() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	})
	return makeHTTPServerWithMux(mux)
}
