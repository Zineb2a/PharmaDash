package util

import db "pharmaDashServer/db/sqlc"

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
