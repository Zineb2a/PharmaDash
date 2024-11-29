package server

import (
	"context"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"

	//"pharmaDashServer/util"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) AddFeedback(c *gin.Context) {
	var req struct {
		OrderID int32  `json:"order_id"`
		Rating  int32  `json:"rating"`
		Comment string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid request."})
		return
	}

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

	// 1. Validate that the order is "Delivered"
	orderStatus, err := query.GetOrderStatus(ctx, req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to get order status."})
		return
	}
	if !orderStatus.Valid || orderStatus.String != "Delivered" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Feedback can only be left for delivered orders."})
		return
	}

	// 2. Retrieve client details
	client, err := query.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to get user."})
		return
	}

	// 3. Add feedback to the database
	if req.Rating > 5 || req.Rating <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid rating. Rating must be between 0 and 5 inclusively."})
		return
	}

	feedback, err := query.AddFeedback(ctx, db.AddFeedbackParams{
		OrderID:  req.OrderID,
		ClientID: client.AccountID,
		Rating:   req.Rating,
		Comment:  pgtype.Text{String: req.Comment, Valid: true}, //had error for type mismatch
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to add feedback.", "error": err.Error()})
		return
	}

	// 4. Send acknowledgment email
	// feedbackEmailBody := "Thank you for your feedback!\n\nYour feedback:\n" + "Rating: " + strconv.Itoa(int(req.Rating)) + "\n" + "Comment: " + req.Comment
	// err = util.SendEmail(email, "Thank you for your feedback!", feedbackEmailBody)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to send acknowledgment email."})
	// 	return
	// }

	// 5. Respond with success
	c.JSON(http.StatusOK, gin.H{"status": "Feedback added successfully.", "feedback": feedback})
}
