package storer

import (
	"kube-watcher/config"
	"kube-watcher/handlers/storer/endpoints"
	"sync"

	"github.com/Sirupsen/logrus"
	api_v1 "k8s.io/api/core/v1"
)

type Storer struct {
	enpoints sync.Map
	logger   *logrus.Entry
}

func (d *Storer) Init(c *config.Config) error {
	d.logger = logrus.WithField("app", "kube-watcher")
	return nil
}

func (d *Storer) ObjectCreated(obj interface{}) {
	d.logger.Info("[Storer Handler] ObjectCreated")
	switch kobj := obj.(type) {
	case *api_v1.Endpoints:
		d.enpoints.Store(endpoints.Transform(kobj))
	}
}

func (d *Storer) ObjectDeleted(obj interface{}) {
	d.logger.Info("[Storer Handler] ObjectDeleted")
	switch ep := obj.(type) {
	case *api_v1.Endpoints:
		d.showIPs(ep)
	}
}

func (d *Storer) ObjectUpdated(obj interface{}) {
	d.logger.Info("[Storer Handler] ObjectUpdated")
	switch ep := obj.(type) {
	case *api_v1.Endpoints:
		d.showIPs(ep)
	}
}

func (d *Storer) showIPs(ep *api_v1.Endpoints) {
	for _, sub := range ep.Subsets {
		for _, ip := range sub.Addresses {
			d.logger.Infof("%s %s %s", ep.Name, ep.Namespace, ip.IP)
		}
	}
}
