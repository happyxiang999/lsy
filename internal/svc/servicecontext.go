package svc

import (
	"lsy/internal/config"
	"lsy/internal/store"
)

type ServiceContext struct {
	Config config.Config
	Store  *store.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Store:  store.NewStore(),
	}
}
