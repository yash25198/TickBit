package server

import (
	"github.com/crema-labs/sxg-go/internal/handler"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	priv_key string
	logger   *zap.Logger
}

func NewServer(priv_key string, logger *zap.Logger) *Server {
	r := gin.Default()
	s := &Server{router: r, priv_key: priv_key, logger: logger}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	hpr := handler.HandleProofRequest{
		PrivKey: s.priv_key,
		Logger:  s.logger,
	}
	s.router.POST("/proof", hpr.HandleProofRequest)
	s.router.GET("/status", hpr.HandleGetStatus)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
