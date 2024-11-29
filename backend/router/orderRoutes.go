package server

import (
	"context"
	"encoding/json"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) CreateOrder(c *gin.Context) {
	payload := c.MustGet("auth_payload").(*token.Payload)
	email := payload.Username
	var req util.OrderRequest
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid request."})
		return
	}

	// Validate card information
	if err := util.ValidateCardInfo(req.CardNumber, req.ExpiryDate, req.CVV); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	ctx := context.Background()
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Rollback(ctx)
	query := db.New(conn)

	dbUser, err := query.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}

	// Retrieve cart and quotation info
	cart, err := query.GetShoppingCartByClientID(ctx, dbUser.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Cart not found."})
		return
	}

	quotation, err := query.GetQuotationByCartID(ctx, pgtype.Int4{Int32: cart.CartID, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Quotation not found."})
		return
	}

	// Create an order from the cart
	orderID, err := query.CreateOrder(ctx, db.CreateOrderParams{
		AccountID:         cart.AccountID,
		DeliveryFrequency: quotation.DeliveryFrequency,
		Destination:       quotation.Destination,
		SpecialHandling:   quotation.SpecialHandling,
		IncludeInsurance:  quotation.IncludeInsurance,
		TotalCost:         quotation.TotalCost,
		Insurance:         quotation.Insurance,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to create order.", "error": err.Error()})
		return
	}

	// Move items from cart to order items
	cartItems, err := query.GetAllShoppingCartItems(ctx, cart.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to retrieve cart items."})
		return
	}

	for _, item := range cartItems {
		_, err := query.CreateOrderItem(ctx, db.CreateOrderItemParams{
			OrderID:         orderID,
			InventoryItemID: item.InventoryItemID,
			Quantity:        pgtype.Int4{Int32: 1, Valid: true}, // Wrapping 1 in pgtype.Int4
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to create order item."})
			return
		}
	}

	// Delete the cart and its items
	_, err = query.DeleteAllCartItems(ctx, cart.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to delete cart items."})
		return
	}

	// Delete the quotation
	err = query.DeleteQuotation(ctx, quotation.QuotationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to delete quotation."})
		return
	}

	_, err = query.DeleteCart(ctx, cart.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to delete cart."})
		return
	}

	// Commit the transaction
	if err := conn.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Transaction failed."})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"status": "Order created successfully", "order_id": orderID})
}

func (server *Server) GetAllOrdersClient(c *gin.Context) {
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

	dbOrders, err := query.GetAllClientOrders(ctx, dbUser.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": dbOrders})
}

func (server *Server) GetAllOrdersAdmin(c *gin.Context) {
	ctx := context.Background()
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error.", "error": err.Error()})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	dbOrders, err := query.GetAllOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": dbOrders})
}
