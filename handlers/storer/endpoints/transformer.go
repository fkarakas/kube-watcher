package endpoints

import (
	api_v1 "k8s.io/api/core/v1"
	"fmt"
)

type Endpoints struct {
	Name      string
	Namespace string
	IPs       []string
}

func Transform(ep *api_v1.Endpoints) (string, Endpoints) {
	key := fmt.Sprintf("%s.%s", ep.Name, ep.Namespace)

	value := Endpoints{
		Name:      ep.Name,
		Namespace: ep.Namespace,
	}

	for _, sub := range ep.Subsets {
		for _, ip := range sub.Addresses {
			//d.logger.Infof("%s %s %s", ep.Name, ep.Namespace, ip.IP)
			value.IPs = append(value.IPs, ip.IP)
		}
	}

	return key, value
}
