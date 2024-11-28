package server

import (
	"context"
	"database/sql"

	//"encoding/json"
	//"io"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"

	//"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	//"github.com/shopspring/decimal"
	"math/big"
)

type QuotationResponse struct {
	Status      string  `json:"status"`
	TotalCost   float64 `json:"total_cost"`
	Insurance   float64 `json:"insurance_cost,omitempty"`
	DeliveryFee float64 `json:"delivery_fee"`
}

// Contains specifications of the delivery service
type QuotationRequest struct {
	QuotationID       int32          `json:"quotation_id"` // to keep track of which quotation was reufsed or accepted
	TotalCost         float64        `json:"total_cost"`
	DeliveryFrequency pgtype.Text    `json:"delivery_frequency"`
	Destination       pgtype.Text    `json:"destination"`
	SpecialHandling   pgtype.Text    `json:"special_handling,omitempty"`
	Insurance         pgtype.Numeric `json:"insurance"`
	IncludeInsurance  bool           `json:"include_insurance"`
	CartID            int32          `json:"cart_id"`
}

func (server *Server) CreateDeliveryQuotation(c *gin.Context) {
	var req QuotationRequest // stores quotation request details
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

	// Check for existing quotation for the cart
	// If a quotation exists for this cart, delete it as a new one is being generated
	existingQuotation, err := query.GetQuotationByCartID(ctx, pgtype.Int4{Int32: req.CartID})

	if err != nil {
		if err == sql.ErrNoRows {
			// No existing quotation found, proceed to create a new one
		} else {
			// Other database error occurred
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to fetch existing quotation."})
			return
		}
	} else {
		if existingQuotation.QuotationID != 0 {
			// Existing quotation found, delete it as a new one is being generated
			err = query.DeleteQuotationByCartID(ctx, pgtype.Int4{Int32: req.CartID})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to replace existing quotation."})
				return
			}
		}
	}

	totalCost := 0.0
	for _, item := range cartItems {
		price := item.UnitPrice
		totalCost += price
	}

	insuranceRatio := 0.1                    // 10% of total cost
	includeInsurance := req.IncludeInsurance // boolean
	deliveryFee := 5.0
	insuranceCost := 0.0 // fixed delivery fee

	if includeInsurance {
		insuranceCost = totalCost * insuranceRatio
		totalCost -= insuranceCost
	}

	totalCost += deliveryFee

	pgNumericInsurance := createPgNumeric(insuranceCost)
	pgNumericTotalCost := createPgNumeric(totalCost)

	createQuotationParams := db.CreateQuotationParams{
		TotalCost:         pgNumericTotalCost,
		DeliveryFrequency: req.DeliveryFrequency,
		Destination:       req.Destination,
		SpecialHandling:   req.SpecialHandling,
		Insurance:         pgNumericInsurance,
		IncludeInsurance:  pgtype.Bool{Bool: true, Valid: true},
	}

	// Call CreateQuotation
	createdQuotation, err := query.CreateQuotation(ctx, createQuotationParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to create quotation."})
		return
	}
	// Type assertion for TotalCost
	totalCost, acc := createdQuotation.TotalCost.Int.Float64()
	if acc != big.Exact {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid TotalCost type."})
		return
	}

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

// helper function
func createPgNumeric(value float64) pgtype.Numeric {
	bigInt := new(big.Int)
	bigInt.SetInt64(int64(value * 100)) // Convert to cents to avoid float precision issues
	return pgtype.Numeric{Int: bigInt, Valid: true}
}

func (server *Server) DeleteQuotation(c *gin.Context) {
	var req struct {
		QuotationID int32 `json:"quotation_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid request"})
		return
	}

	ctx := context.Background()
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error"})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	// Execute delete query
	err = query.DeleteQuotation(ctx, req.QuotationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to delete quotation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Quotation deleted successfully"})
}
