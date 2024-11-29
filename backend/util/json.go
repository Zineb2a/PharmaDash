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
	CartID      int32 `json:"cart_id"`
	InventoryID int32 `json:"inventory_id"`
}

type RemoveItemFromCartRequest struct {
	CartItemID      int32 `json:"cart_item_id"`
	InventoryItemID int32 `json:"inventory_item_id"`
}

type DeleteCartRequest struct {
	CartID int32
}

type OrderRequest struct {
	CartID     int32  `json:"cart_id"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
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
	PharmacyID       int32   `json:"pharmacy_id"`
	Item_Name        string  `json:"item_name"`
	Item_Description string  `json:"item_description"`
	Medication_Name  string  `json:"medication_name"`
	Unit_price       float32 `json:"unit_price"`
	Stock_Quantity   int32   `json:"stock_quantity"`
	OTC              bool    `json:"otc"`
}
