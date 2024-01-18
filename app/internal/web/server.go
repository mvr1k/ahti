package web

import (
	"ahti/app/config"
	"ahti/app/internal/reminders"
	"ahti/app/internal/web/api"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	router chi.Router
	apis   map[string]api.API //module name to API mapping
	logger *zap.SugaredLogger
}

const version = "/api/v1"

func NewServer(logger *zap.SugaredLogger) *Server {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	return &Server{
		router: router,
		apis:   map[string]api.API{},
		logger: logger,
	}
}

func (server *Server) registerModuleToMap(module api.API) {
	server.apis[module.ModuleName()] = module
}

func (server *Server) Start() (err error) {
	address := fmt.Sprintf("%v:%v", config.ServerHost, config.ServerPort)

	//add all modeules one by one
	server.registerModuleToMap(reminders.NewApiController(server.logger))

	server.addAllRoutesToServer()
	server.logger.Infof("Starting Server at %v", address)
	err = http.ListenAndServe(address, server.router)
	return
}

func (server *Server) addAllRoutesToServer() {
	server.router.Use(middleware.DefaultLogger)
	for _, apiModule := range server.apis {
		for _, route := range apiModule.GetRoutingList() {
			server.router.MethodFunc(route.Method, version+"/"+route.Path, route.Handler)
		}
	}
}
