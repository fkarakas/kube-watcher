package handlers

import (
	"kube-watcher/config"
)

type Handler interface {
	Init(c *config.Config) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(obj interface{})
}
