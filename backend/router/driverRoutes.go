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

func (server *Server) ConfirmOrderDelivery(c *gin.Context) {
	// Get the driver's email from the token payload
	payload := c.MustGet("auth_payload").(*token.Payload)
	if payload.Role != "Driver" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid Request: only drivers can perform this operation"})
		return
	}

	//driverEmail := payload.Username
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
	err = util.SendEmail(userEmail.String, "Your order has been delivered", "Your order has been successfully delivered by driver: "+payload.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to send confirmation email."})
		return
	}

	// Commit the transaction
	err = conn.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to commit transaction."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order delivery confirmed successfully."})
}

func (server *Server) PickUpOrder(c *gin.Context) {
	// Get the driver's email from the token payload
	payload := c.MustGet("auth_payload").(*token.Payload)
	if payload.Role != "Driver" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid Request: only drivers can perform this operation"})
		return
	}
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

	//id := ConvertUUIDToPGTypeUUID(payload.ID)
	assignOrder := db.AssignOrderToDriverParams{
		DriverID: pgtype.Int4{Int32: 1, Valid: true},
		OrderID:  req.OrderID,
	}

	// Lock the order for this driver and update its status
	if err := query.AssignOrderToDriver(ctx, assignOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to assign order."})
		return
	}

	// Get the user's email associated with the order
	userEmail, err := query.GetEmailFromOrder(ctx, selectedOrder.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to retrieve email."})
		return
	}

	//Send email notification to the user
	err = util.SendEmail(userEmail.String, "Your order is out for delivery", "Your order is now being delivered by driver: "+driver.Name)
	if err != nil {
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

func (server *Server) ReadyOrdersForPickup(c *gin.Context) {
	// Extract authentication payload
	payload := c.MustGet("auth_payload").(*token.Payload)
	if payload.Role != "Driver" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only drivers can perform this operation"})
		return
	}

	// Initialize context and database connection
	ctx := context.Background()
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start database transaction"})
		return
	}
	defer conn.Rollback(ctx) // Ensure rollback on function exit

	// Fetch available orders
	query := db.New(conn)
	orders, err := query.GetAvailableOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available orders"})
		return
	}

	// Commit the transaction
	if err := conn.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Respond with the fetched orders
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
