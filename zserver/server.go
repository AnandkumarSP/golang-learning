package zserver

import (
	"net/http"

	"../logger"
	"../utils"
)

// ZServer definition
type ZServer struct {
	server *http.Server
}

// New creates a ZServer instance and return its pointer
func New() (s *ZServer) {
	s = &ZServer{
		server: &http.Server{
			Addr: ":8080",
		},
	}
	s.init()
	return
}

func (s *ZServer) init() {
	logger.Info("Registering routes...")
	s.registerRoutes()
}

// Start runs the ZServer starting sequence
func (s *ZServer) Start() {
	logger.Info("Reading uploader configuration")
	utils.RefreshUploaderConfig()
	logger.Info("Reading plugins configuration")
	utils.RefreshPluginsConfig()
	logger.Infof("Starting server at port '%s'...", s.server.Addr)
	s.server.ListenAndServe()
}
