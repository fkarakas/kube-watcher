package main

import (
	"kube-watcher/config"
	"kube-watcher/controller"
	"kube-watcher/handlers"
	"kube-watcher/handlers/storer"
	"log"
)

func main() {
	var eventHandler handlers.Handler
	eventHandler = new(storer.Storer)

	conf := &config.Config{
		Resource: config.Resource{
			Endpoints: true,
		},
		Handler:   config.Handler{},
		Namespace: "fatih",
	}

	if err := eventHandler.Init(conf); err != nil {
		log.Fatal(err)
	}

	controller.Start(conf, eventHandler)
}
