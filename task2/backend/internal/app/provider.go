package app

import (
	"log"
	"os"

	nethttp "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"software-engineering-2/internal/config"
	_map "software-engineering-2/internal/delivery/map"
	"software-engineering-2/internal/infrastructure/http"
	_map3 "software-engineering-2/internal/storage/map"
	storageGeneral "software-engineering-2/internal/storage/map/general"
	_map2 "software-engineering-2/internal/usecase/map"
	usecaseGeneral "software-engineering-2/internal/usecase/map/general"
)

type provider struct {
	config *config.Config

	server  *http.Server
	handler nethttp.Handler

	mapDelivery *_map.Delivery

	mapUseCase _map2.UseCase

	mapStorage _map3.Storage
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) Config() *config.Config {
	if p.config == nil {
		cfg, err := config.NewConfig(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		p.config = cfg
	}

	return p.config
}

func (p *provider) Server() *http.Server {
	if p.server == nil {
		p.server = http.NewServer(p.Config().HTTPConfig, p.Handler())
	}

	return p.server
}

func (p *provider) Handler() nethttp.Handler {
	if p.handler == nil {
		handler := echo.New()

		handler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{nethttp.MethodGet, nethttp.MethodPut, nethttp.MethodPost, nethttp.MethodPatch, nethttp.MethodDelete},
		}))
		handler.Use(middleware.Logger())
		handler.Use(middleware.Recover())

		apiGroup := handler.Group("/api")
		p.MapDelivery().RegisterRoutes(apiGroup)

		p.handler = handler
	}

	return p.handler
}

func (p *provider) MapDelivery() *_map.Delivery {
	if p.mapDelivery == nil {
		p.mapDelivery = _map.NewMapDelivery(p.MapUseCase())
	}

	return p.mapDelivery
}

func (p *provider) MapUseCase() _map2.UseCase {
	if p.mapUseCase == nil {
		p.mapUseCase = usecaseGeneral.NewUseCase(p.MapStorage())
	}

	return p.mapUseCase
}

func (p *provider) MapStorage() _map3.Storage {
	if p.mapStorage == nil {
		p.mapStorage = storageGeneral.NewStorage()
	}

	return p.mapStorage
}
