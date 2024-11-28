package util

import (
	db "pharmaDashServer/db/sqlc"
)

type RegisterRequest struct {
	RegisterMode string
	UserData     *db.Account
}

type LoginRequest struct {
	Email    string
	Password string
}

type AddItemToCartRequest struct {
	CartID      int32
	InventoryID int32
}

type RemoveItemFromCartRequest struct {
	CartItemID      int32
	InventoryItemID int32
}

type DeleteCartRequest struct {
	CartID int32
}

type OrderRequest struct {
	CartID     int32
	CardNumber string
	ExpiryDate string
	CVV        string
}

type PickUpOrderRequest struct {
	OrderID int32 `json:"order_id" binding:"required"`
}

type ConfirmDeliveryRequest struct {
	OrderID int32 `json:"order_id" binding:"required"`
}

type AddFeedbackRequest struct {
	OrderID int32  `json:"order_id" binding:"required"`
	Rating  int32  `json:"rating" binding:"required"`
	Comment string `json:"comment"`
}

type AddItemToInventory struct {
	PharmacyID       int32
	Item_Name        string
	Item_Description string
	Medication_Name  string
	Unit_price       float32
	Stock_Quantity   int32
	OTC              bool
}
