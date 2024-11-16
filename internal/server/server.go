package server

import (
	"github.com/crema-labs/sxg-go/internal/handler"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	logger  *zap.Logger
	handler *handler.HandleProofRequest
}

func NewServer(priv_key string, tbClient ethereum.TBClient, logger *zap.Logger) *Server {
	r := gin.Default()
	s := &Server{router: r, logger: logger, handler: &handler.HandleProofRequest{
		TBClient: tbClient,
		PrivKey:  priv_key,
		Logger:   logger,
	}}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {

	s.router.POST("/deets", s.handler.HandleDeets)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
