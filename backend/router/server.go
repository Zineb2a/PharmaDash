package server

import (
	"context"
	"net/http"
	"os"
	"pharmaDashServer/token"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Server struct {
	Router *gin.Engine
	maker  *token.PasetoMaker
	pool   *pgxpool.Pool
}

func GetNewServer() (*Server, error) {
	//load env vars
	godotenv.Load()
	key := os.Getenv("TOKEN_KEY")
	tokenMaker, err := token.NewPasetoMaker(key)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	server := &Server{
		Router: router,
		maker:  tokenMaker,
		pool:   pool,
	}

	//mount routes here
	api := router.Group("/api")
	{
		// Gives information about the API in general, particularly about how to switch between versions
		api.GET("", server.NotImplemented)
	}
	// load router
	// load token maker
	return server, nil
}

func (server *Server) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "This resource is not yet implemented, but will be in the future"})
}
