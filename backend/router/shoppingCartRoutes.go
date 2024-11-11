package server

import (
	"context"
	"encoding/json"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) CreateShoppingCart(c *gin.Context) {
	payload := c.MustGet("auth_payload").(*token.Payload)
	email := payload.Username
	ctx := context.Background()

	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	dbUser, err := query.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}

	dbShoppingCart, err := query.GetShoppingCartByClientID(ctx, dbUser.AccountID)
	if err != nil {
		dbShoppingCart, err = query.CreateShoppingCart(ctx, dbUser.AccountID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error, failed to create cart."})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Cart creation successful", "cart_id": dbShoppingCart.CartID})
			return
		}
	} else {
		cartItem, err := query.GetAllShoppingCartItems(ctx, dbShoppingCart.CartID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "Cart already exists", "cart_id": dbShoppingCart.CartID, "cart_items": nil})
			return
		}
		cartItemJSON, err := json.Marshal(cartItem)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "Cart already exists", "cart_id": dbShoppingCart.CartID, "cart_items": nil})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Cart already exists", "cart_id": dbShoppingCart.CartID, "cart_items": cartItemJSON})
		return
	}
}
