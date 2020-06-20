package router

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gabek/owncast/config"
	"github.com/gabek/owncast/controllers"
	"github.com/gabek/owncast/core/chat"
	"github.com/gabek/owncast/core/rtmp"
)

//Start starts the router for the http, ws, and rtmp
func Start() error {
	// websocket server
	chatServer := chat.NewServer("/entry")
	go chatServer.Listen()

	// start the rtmp server
	go rtmp.Start()

	// static files
	http.HandleFunc("/", controllers.IndexHandler)

	// status of the system
	http.HandleFunc("/status", controllers.GetStatus)

	port := config.Config.WebServerPort

	log.Printf("Starting public web server on port: %d", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
