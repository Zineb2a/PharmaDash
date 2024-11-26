package server

import (
	"context"
	"net/http"

	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) ConfirmOrderDelivery(c *gin.Context) {
	// Get the driver's email from the token payload
	payload := c.MustGet("auth_payload").(*token.Payload)
	driverEmail := payload.Username
	ctx := context.Background()

	// Acquire a connection from the pool
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Rollback(ctx)

	query := db.New(conn)

	// Read request body into the new struct
	var req util.ConfirmDeliveryRequest // Define this struct with OrderID and any other necessary fields
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid request."})
		return
	}

	// Update the order status to "completed"
	if _, err := query.MarkOrderAsDelivered(ctx, req.OrderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to update order status."})
		return
	}

	// Get the user's email associated with the order
	userEmail, err := query.GetEmailFromOrder(ctx, req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to retrieve email."})
		return
	}

	// Validate email before sending
	if userEmail.String == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "User  email is not available."})
		return
	}

	// Send email notification to the user
	if err := util.SendEmail(userEmail.String, "Your order has been delivered", "Your order has been successfully delivered by driver: "+driverEmail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to send confirmation email."})
		return
	}

	// Commit the transaction
	if err := conn.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to commit transaction."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order delivery confirmed successfully."})
}
