package endpoints

import (
	"fmt"
	api_v1 "k8s.io/api/core/v1"
)

type Endpoints struct {
	Name      string
	Namespace string
	Subsets   []string
}

func Build(ep *api_v1.Endpoints) (string, Endpoints) {
	key := fmt.Sprintf("%s.%s", ep.Name, ep.Namespace)

	value := Endpoints{
		Name:      ep.Name,
		Namespace: ep.Namespace,
	}

	for _, sub := range ep.Subsets {
		for _, ip := range sub.Addresses {
			//d.logger.Infof("%s %s %s", ep.Name, ep.Namespace, ip.IP)
			value.Subsets = append(value.Subsets, ip.IP)
		}
	}

	return key, value
}
