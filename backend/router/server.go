package server

import (
	"context"
	"net/http"
	"os"
	appsystem "pharmaDashServer/app_system"
	"pharmaDashServer/token"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Server struct {
	Router *gin.Engine
	maker  *token.PasetoMaker
	pool   *pgxpool.Pool
	system *appsystem.System
}

func GetNewServer() (*Server, error) {
	//load env vars
	godotenv.Load()
	key := os.Getenv("TOKEN_KEY")
	tokenMaker, err := token.NewPasetoMaker(key)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	system := appsystem.InitializeSystem()
	server := &Server{
		Router: router,
		maker:  tokenMaker,
		pool:   pool,
		system: system,
	}

	//mount routes here
	api := router.Group("/api")
	{
		api.GET("", server.NotImplemented)
		user := router.Group("/user")
		{
			user.POST("/register", server.RegisterNewAccount)
			user.POST("/login", server.Login)
			user.POST("/logout", server.LogOut)
		}
		cart := router.Group("/cart")
		{
			cart.POST("/create_cart", server.mustAuthChecker, server.CreateShoppingCart)
			cart.POST("/add_cart_item", server.mustAuthChecker, server.addCartItem)
			cart.POST("/remove_cart_item", server.mustAuthChecker, server.removeCartItem)
			cart.POST("/cancel_cart", server.mustAuthChecker, server.cancelShoppingCart)
			cart.POST("/delivery_quotation", server.mustAuthChecker, server.CreateDeliveryQuotation)
			cart.POST("/delivery_quotation/accept", server.mustAuthChecker, server.AcceptQuotation)
			cart.POST("/delivery_quotation/refuse", server.mustAuthChecker, server.RefuseQuotation)
		}
	}
	return server, nil
}

func (server *Server) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "This resource is not yet implemented, but will be in the future"})
}
