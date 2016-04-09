package main

import (
	"flag"
	"fmt"
	"github.com/gesellix/go-webhook"
	"github.com/gesellix/go-webhook/docker-hub"
	"github.com/gesellix/go-webhook/drone"
	"log"
	"net/http"
)

var (
	listenAddress = flag.String("listen-address", "127.0.0.1:3003", "<address>:<port> to listen on")
	configFile    = flag.String("config-file", "config.json", "Location of config file")
)

func main() {
	flag.Parse()

	config, err := webhook.ReadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("using config from %s", *configFile)

	for _, handlerConfig := range config.Handlers {
		newHandler(handlerConfig)
	}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "OK")
	//})

	log.Printf("starting webhook at %s", *listenAddress)

	if config.Tls.Key != "" && config.Tls.Cert != "" {
		log.Print("Starting with SSL")
		http.ListenAndServeTLS(*listenAddress, config.Tls.Cert, config.Tls.Key, Log(http.DefaultServeMux))
	} else {
		log.Print("Warning: Server is starting without SSL, you should not pass any credentials using this configuration")
		log.Print("To use SSL, you must provide a config file with a [tls] section, and provide locations to a `key` file and a `cert` file")
		http.ListenAndServe(*listenAddress, Log(http.DefaultServeMux))
	}
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.RemoteAddr, r.Method)
		handler.ServeHTTP(w, r)
	})
}

func newHandler(h webhook.Handler) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK (default handler)")
	}

	switch h.Type {
	case webhook.DockerHub:
		httpHandler = dockerhub.NewHandler()
		break
	case webhook.Drone:
		httpHandler = drone.NewHandler()
		break
	}

	if h.ApiKey != "" {
		authHttpHandler := func(w http.ResponseWriter, r *http.Request) {
			apikey := r.URL.Query().Get("apikey")
			//apikey := r.Header.Get("Authorization")
			if h.ApiKey != apikey {
				http.Error(w, "Not Authorized", 401)
			} else {
				httpHandler(w, r)
			}
		}
		log.Printf("installing handler at %q", h.Path)
		http.HandleFunc(h.Path, authHttpHandler)
	} else {
		log.Printf("Warning: The handler for type %q at %q is about to start without any authentication.", h.Type, h.Path)
		log.Println("Anyone can trigger handlers to fire off.")
		log.Println("To enable authentication, you must add an `apikey` attribute.")
		log.Printf("installing handler at %q", h.Path)
		http.HandleFunc(h.Path, httpHandler)
	}
}
