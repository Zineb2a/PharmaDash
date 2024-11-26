package server

import (
	"context"
	"net/http"

	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) PickUpOrder(c *gin.Context) {
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

	// Get driver details
	driver, err := query.GetDriverByEmail(ctx, pgtype.Text{String: driverEmail, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Driver not found."})
		return
	}

	// Read request body into the new struct
	var req util.PickUpOrderRequest // Using the new struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid request."})
		return
	}

	// Fetch available orders
	availableOrders, err := query.GetAvailableOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to fetch available orders."})
		return
	}

	// Check if the requested order is available
	var selectedOrder db.Order
	for _, order := range availableOrders {
		if order.OrderID == req.OrderID {
			selectedOrder = order
			break
		}
	}

	if selectedOrder.OrderID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Order not available."})
		return
	}

	// Lock the order for this driver and update its status
	if _, err := query.AssignOrderToDriver(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to assign order."})
		return
	}

	// Get the user's email associated with the order
	userEmail, err := query.GetEmailFromOrder(ctx, selectedOrder.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to retrieve email."})
		return
	}

	// Send email notification to the user
	if err := util.SendEmail(userEmail.String, "Your order is out for delivery", "Your order is now being delivered by driver: "+driver.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to send notification email."})
		return
	}
	// Commit the transaction
	if err := conn.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to commit transaction."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order picked up successfully."})
}
