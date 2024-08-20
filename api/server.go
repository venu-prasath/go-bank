package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/venu-prasath/go-bank/db/sqlc"
)

type Server struct {
	store db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.GetAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}