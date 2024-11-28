package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	db "pharmaDashServer/db/sqlc"
	"pharmaDashServer/token"
	"pharmaDashServer/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (server *Server) CreateShoppingCart(c *gin.Context) {
	payload := c.MustGet("auth_payload").(*token.Payload)
	if payload.Role != "Client" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid Request: only clients can perform this operation"})
		return
	}
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

func (server *Server) addCartItem(c *gin.Context) {
	ctx := context.Background()

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	var payload util.AddItemToCartRequest
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid json payload format."})
		return
	}

	//acquire connection from connection pools
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Rollback(ctx)
	query := db.New(nil)
	query = query.WithTx(conn)

	//Inventory Id
	inventoryID := payload.InventoryID

	//find available item
	inventoryItem, err := query.GetAvailableInventoryItem(ctx, inventoryID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Item out of stock."})
		return
	}

	//mark reserved
	_, err = query.ReserveItem(ctx, inventoryItem.InventoryID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to reserve item."})
		return
	}

	//create shopping cart Item
	params2 := db.CreateShoppingCartItemParams{
		CartID:          payload.CartID,
		InventoryItemID: inventoryItem.InventoryID,
	}
	cartItem, err := query.CreateShoppingCartItem(ctx, params2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to create cart item."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Cart item added successfully", "CartItemID": cartItem.ShoppingCartItemID, "InventoryItemID": inventoryItem.InventoryID})
	conn.Commit(ctx)
}

func (server *Server) removeCartItem(c *gin.Context) {
	ctx := context.Background()

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	var payload util.RemoveItemFromCartRequest
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid json payload format."})
		return
	}

	//acquire connection from connection pools
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Rollback(ctx)
	query := db.New(nil)
	query = query.WithTx(conn)

	//mark reserved
	_, err = query.FreeItem(ctx, payload.InventoryItemID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to reserve item."})
		return
	}

	//Delete cart item
	_, err = query.DeleteCartItem(ctx, payload.CartItemID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to delete cart item."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Cart item deleted successfully"})
	conn.Commit(ctx)
}

func (server *Server) cancelShoppingCart(c *gin.Context) {
	ctx := context.Background()

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	var payload util.DeleteCartRequest
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Invalid json payload format."})
		return
	}

	//acquire connection from connection pools
	conn, err := server.pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Rollback(ctx)
	query := db.New(nil)
	query = query.WithTx(conn)

	//mark all items as available
	_, err = query.FreeAllCartItems(ctx, payload.CartID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to free inventory items."})
		return
	}
	//increment stock in inventory

	//delete all cart items DONE
	_, err = query.DeleteAllCartItems(ctx, payload.CartID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to delete cart items."})
		return
	}
	//delete cart
	_, err = query.DeleteCart(ctx, payload.CartID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed to delete cart."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Cart deleted successfully"})
	conn.Commit(ctx)
}

func (server *Server) getAllItems(c *gin.Context) {
	ctx := context.Background()
	conn, err := server.pool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Server error."})
		return
	}
	defer conn.Release()
	query := db.New(conn)

	dbInventory, err := query.GetAllInventoryItems(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Database error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Items": dbInventory})
}

func (server *Server) addItemToInventory(c *gin.Context) {

}
