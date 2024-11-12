package server

import (
	"context"
	//"encoding/json"
	//"io"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"

	//"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type QuotationResponse struct {
	Status      string  `json:"status"`
	TotalCost   float64 `json:"total_cost"`
	Insurance   float64 `json:"insurance_cost,omitempty"`
	DeliveryFee float64 `json:"delivery_fee"`
}

func (server *Server) CreateDeliveryQuotation(c *gin.Context) {

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

	// to get the shopping cart
	dbShoppingCart, err := query.GetShoppingCartByClientID(ctx, dbUser.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Cart not found."})
		return
	}

	// extract items from shopping cart
	cartItems, err := query.GetAllShoppingCartItems(ctx, dbShoppingCart.CartID)
	if err != nil || len(cartItems) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "Cart is empty. No cart items found."})
		return
	}

	totalCost := 0.0
	for _, item := range cartItems {
		price, _ := item.UnitPrice.Int.Float64()
		totalCost += price
	}

	insuranceCost := 10.0
	includeInsurance := true
	deliveryFee := 5.0 // fixed delivery fee

	if includeInsurance {
		totalCost += insuranceCost
	}

	totalCost += deliveryFee

	// Set the struct
	quotation := QuotationResponse{
		Status:      "Quotation generated",
		TotalCost:   totalCost,
		Insurance:   insuranceCost,
		DeliveryFee: deliveryFee,
	}

	// Return the quotation to the user
	c.JSON(http.StatusOK, quotation)

}
