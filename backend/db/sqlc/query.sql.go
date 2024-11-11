// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createShoppingCart = `-- name: CreateShoppingCart :one
INSERT INTO ShoppingCart (account_id) 
VALUES ($1) RETURNING cart_id, account_id
`

func (q *Queries) CreateShoppingCart(ctx context.Context, accountID int32) (Shoppingcart, error) {
	row := q.db.QueryRow(ctx, createShoppingCart, accountID)
	var i Shoppingcart
	err := row.Scan(&i.CartID, &i.AccountID)
	return i, err
}

const createShoppingCartItem = `-- name: CreateShoppingCartItem :one
INSERT INTO ShoppingCartItems (cart_id, inventory_item_id) 
VALUES ($1,$2) RETURNING shopping_cart_item_id, cart_id, inventory_item_id
`

type CreateShoppingCartItemParams struct {
	CartID          int32
	InventoryItemID int32
}

func (q *Queries) CreateShoppingCartItem(ctx context.Context, arg CreateShoppingCartItemParams) (Shoppingcartitem, error) {
	row := q.db.QueryRow(ctx, createShoppingCartItem, arg.CartID, arg.InventoryItemID)
	var i Shoppingcartitem
	err := row.Scan(&i.ShoppingCartItemID, &i.CartID, &i.InventoryItemID)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO Accounts (name, last_name, password, phone_number, email, address) 
VALUES ($1,$2,$3,$4,$5,$6) RETURNING account_id, name, last_name, password, phone_number, email, address, authlevel
`

type CreateUserParams struct {
	Name        string
	LastName    string
	Password    string
	PhoneNumber pgtype.Text
	Email       pgtype.Text
	Address     pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Account, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.LastName,
		arg.Password,
		arg.PhoneNumber,
		arg.Email,
		arg.Address,
	)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Name,
		&i.LastName,
		&i.Password,
		&i.PhoneNumber,
		&i.Email,
		&i.Address,
		&i.Authlevel,
	)
	return i, err
}

const decrementInventoryStock = `-- name: DecrementInventoryStock :one
UPDATE Inventory SET stock = stock - 1 
WHERE inventory_id = $1 RETURNING inventory_id, pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc
`

func (q *Queries) DecrementInventoryStock(ctx context.Context, inventoryID int32) (Inventory, error) {
	row := q.db.QueryRow(ctx, decrementInventoryStock, inventoryID)
	var i Inventory
	err := row.Scan(
		&i.InventoryID,
		&i.PharmacyID,
		&i.ItemName,
		&i.ItemDescription,
		&i.MedicationName,
		&i.UnitPrice,
		&i.StockQuantity,
		&i.Otc,
	)
	return i, err
}

const deleteCartItem = `-- name: DeleteCartItem :one
DELETE FROM ShoppingCartItems 
WHERE shopping_cart_item_id = $1 RETURNING shopping_cart_item_id, cart_id, inventory_item_id
`

func (q *Queries) DeleteCartItem(ctx context.Context, shoppingCartItemID int32) (Shoppingcartitem, error) {
	row := q.db.QueryRow(ctx, deleteCartItem, shoppingCartItemID)
	var i Shoppingcartitem
	err := row.Scan(&i.ShoppingCartItemID, &i.CartID, &i.InventoryItemID)
	return i, err
}

const freeItem = `-- name: FreeItem :one
UPDATE InventoryItems SET reserved = 0
WHERE inventory_item_id = $1 RETURNING inventory_item_id, inventory_id, reserved
`

func (q *Queries) FreeItem(ctx context.Context, inventoryItemID int32) (Inventoryitem, error) {
	row := q.db.QueryRow(ctx, freeItem, inventoryItemID)
	var i Inventoryitem
	err := row.Scan(&i.InventoryItemID, &i.InventoryID, &i.Reserved)
	return i, err
}

const getAllShoppingCartItems = `-- name: GetAllShoppingCartItems :many
SELECT shopping_cart_item_id, cart_id, sci.inventory_item_id, ii.inventory_item_id, ii.inventory_id, reserved, i.inventory_id, pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc
FROM ShoppingCartItems AS sci
INNER JOIN InventoryItems AS ii ON sci.inventory_item_id = ii.inventory_item_id
INNER JOIN Inventory AS i ON ii.inventory_id = i.inventory_id
WHERE sci.cart_id = $1
`

type GetAllShoppingCartItemsRow struct {
	ShoppingCartItemID int32
	CartID             int32
	InventoryItemID    int32
	InventoryItemID_2  int32
	InventoryID        int32
	Reserved           pgtype.Bits
	InventoryID_2      int32
	PharmacyID         int32
	ItemName           string
	ItemDescription    string
	MedicationName     string
	UnitPrice          pgtype.Numeric
	StockQuantity      int32
	Otc                pgtype.Bits
}

func (q *Queries) GetAllShoppingCartItems(ctx context.Context, cartID int32) ([]GetAllShoppingCartItemsRow, error) {
	rows, err := q.db.Query(ctx, getAllShoppingCartItems, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllShoppingCartItemsRow
	for rows.Next() {
		var i GetAllShoppingCartItemsRow
		if err := rows.Scan(
			&i.ShoppingCartItemID,
			&i.CartID,
			&i.InventoryItemID,
			&i.InventoryItemID_2,
			&i.InventoryID,
			&i.Reserved,
			&i.InventoryID_2,
			&i.PharmacyID,
			&i.ItemName,
			&i.ItemDescription,
			&i.MedicationName,
			&i.UnitPrice,
			&i.StockQuantity,
			&i.Otc,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAvailableInventoryItem = `-- name: GetAvailableInventoryItem :one
SELECT inventory_item_id, inventory_id, reserved FROM InventoryItems
WHERE inventory_item_id = $1
AND reserved = 0 LIMIT 1
`

func (q *Queries) GetAvailableInventoryItem(ctx context.Context, inventoryItemID int32) (Inventoryitem, error) {
	row := q.db.QueryRow(ctx, getAvailableInventoryItem, inventoryItemID)
	var i Inventoryitem
	err := row.Scan(&i.InventoryItemID, &i.InventoryID, &i.Reserved)
	return i, err
}

const getInventoryItemByID = `-- name: GetInventoryItemByID :one
SELECT inventory_item_id, inventory_id, reserved FROM InventoryItems
WHERE inventory_item_id = $1 LIMIT 1
`

func (q *Queries) GetInventoryItemByID(ctx context.Context, inventoryItemID int32) (Inventoryitem, error) {
	row := q.db.QueryRow(ctx, getInventoryItemByID, inventoryItemID)
	var i Inventoryitem
	err := row.Scan(&i.InventoryItemID, &i.InventoryID, &i.Reserved)
	return i, err
}

const getShoppingCartByClientID = `-- name: GetShoppingCartByClientID :one
SELECT cart_id, account_id FROM ShoppingCart
WHERE account_id = $1 LIMIT 1
`

func (q *Queries) GetShoppingCartByClientID(ctx context.Context, accountID int32) (Shoppingcart, error) {
	row := q.db.QueryRow(ctx, getShoppingCartByClientID, accountID)
	var i Shoppingcart
	err := row.Scan(&i.CartID, &i.AccountID)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT account_id, name, last_name, password, phone_number, email, address, authlevel FROM Accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (Account, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Name,
		&i.LastName,
		&i.Password,
		&i.PhoneNumber,
		&i.Email,
		&i.Address,
		&i.Authlevel,
	)
	return i, err
}

const incrementInventoryStock = `-- name: IncrementInventoryStock :one
UPDATE Inventory SET stock = stock - 1 
WHERE inventory_id = $1 RETURNING inventory_id, pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc
`

func (q *Queries) IncrementInventoryStock(ctx context.Context, inventoryID int32) (Inventory, error) {
	row := q.db.QueryRow(ctx, incrementInventoryStock, inventoryID)
	var i Inventory
	err := row.Scan(
		&i.InventoryID,
		&i.PharmacyID,
		&i.ItemName,
		&i.ItemDescription,
		&i.MedicationName,
		&i.UnitPrice,
		&i.StockQuantity,
		&i.Otc,
	)
	return i, err
}

const reserveItem = `-- name: ReserveItem :one
UPDATE InventoryItems SET reserved = 1
WHERE inventory_item_id = $1 RETURNING inventory_item_id, inventory_id, reserved
`

func (q *Queries) ReserveItem(ctx context.Context, inventoryItemID int32) (Inventoryitem, error) {
	row := q.db.QueryRow(ctx, reserveItem, inventoryItemID)
	var i Inventoryitem
	err := row.Scan(&i.InventoryItemID, &i.InventoryID, &i.Reserved)
	return i, err
}
