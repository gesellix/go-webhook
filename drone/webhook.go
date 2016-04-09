package drone

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//go:generate gojson -input example.json -o webhook-gen.go -pkg drone -name Drone

func NewHandler() func(w http.ResponseWriter, r *http.Request) {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Panic(err)
		}
		defer r.Body.Close()

		var message Drone
		if err := json.Unmarshal(body, &message); err != nil {
			log.Printf("got invalid webhook message %q", err)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			fmt.Fprintf(w, "Got an invalid webhook message. Please check the logs.")
		} else {
			log.Printf("got webhook message %v", message)

			backendImage := fmt.Sprintf("foo/backend:%s", message.Docker.Images[0].Tag)
			frontendImage := fmt.Sprintf("foo/frontend:%s", message.Docker.Images[0].Tag)
			log.Printf("going to deploy %q and %q", frontendImage, backendImage)

			// docker pull foo/ansible:latest
			// docker run --rm -it .... foo/ansible deploy -t message.Docker.Images[0].Tag
		}

	}
	return httpHandler
}
