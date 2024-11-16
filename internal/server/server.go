package server

import (
	"github.com/crema-labs/sxg-go/internal/handler"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	s := &Server{router: r, logger: logger, handler: &handler.HandleProofRequest{
		TBClient: tbClient,
		PrivKey:  priv_key,
		Logger:   logger,
	}}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	s.router.GET("/deets", s.handler.HandleDeets)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
