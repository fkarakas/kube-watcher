package main

import (
	"kube-watcher/api"
	"kube-watcher/config"
	"kube-watcher/controller"
	"kube-watcher/handlers/storer"
	"log"
)

func main() {
	//var eventHandler handlers.Handler
	handler := new(storer.Storer)

	conf := &config.Config{
		Resource: config.Resource{
			Endpoints: true,
		},
		Handler:   config.Handler{},
		Namespace: "default",
	}

	if err := handler.Init(conf); err != nil {
		log.Fatal(err)
	}

	api.Start(handler)

	controller.Start(conf, handler)
}
