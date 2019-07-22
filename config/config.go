package config

type Config struct {
	Handler   Handler  `json:"handler"`
	Resource  Resource `json:"resource"`
	Namespace string   `json:"namespace,omitempty"`
}

type Resource struct {
	Endpoints bool `json:"ep"`
}

type Handler struct {
}
