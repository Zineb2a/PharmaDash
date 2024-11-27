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

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
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
			cart.POST("/delete_quotation", server.mustAuthChecker, server.DeleteQuotation)
			// cart.POST("/delivery_quotation/accept", server.mustAuthChecker, server.AcceptQuotation)
			// cart.POST("/delivery_quotation/refuse", server.mustAuthChecker, server.RefuseQuotation)

		}
		order := router.Group("/order")
		{
			order.POST("/create_order", server.CreateOrder)
			order.GET("/get_orders", server.mustAuthChecker, server.GetAllOrders)
		}
		feedback := router.Group("/feedback")
		{
			feedback.POST("/feedback", server.mustAuthChecker, server.AddFeedback)
		}
	}
	return server, nil
}

func (server *Server) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "This resource is not yet implemented, but will be in the future"})
}
