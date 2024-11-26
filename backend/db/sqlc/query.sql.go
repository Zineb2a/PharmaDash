// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addFeedback = `-- name: AddFeedback :one
INSERT INTO Feedback (order_id, client_id, rating, comment)
VALUES ($1, $2, $3, $4)
RETURNING feedback_id, order_id, client_id, rating, comment, created_at
`

type AddFeedbackParams struct {
	OrderID  int32
	ClientID int32
	Rating   int32
	Comment  pgtype.Text
}

func (q *Queries) AddFeedback(ctx context.Context, arg AddFeedbackParams) (Feedback, error) {
	row := q.db.QueryRow(ctx, addFeedback,
		arg.OrderID,
		arg.ClientID,
		arg.Rating,
		arg.Comment,
	)
	var i Feedback
	err := row.Scan(
		&i.FeedbackID,
		&i.OrderID,
		&i.ClientID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const assignOrderToDriver = `-- name: AssignOrderToDriver :one


BEGIN
`

type AssignOrderToDriverRow struct {
}

// Order is still in a pending state
func (q *Queries) AssignOrderToDriver(ctx context.Context) (AssignOrderToDriverRow, error) {
	row := q.db.QueryRow(ctx, assignOrderToDriver)
	var i AssignOrderToDriverRow
	err := row.Scan()
	return i, err
}

const createDriver = `-- name: CreateDriver :one
INSERT INTO Drivers (name, email, phone_number)
VALUES ($1, $2, $3)
RETURNING driver_id, name, email, phone_number
`

type CreateDriverParams struct {
	Name        string
	Email       pgtype.Text
	PhoneNumber pgtype.Text
}

type CreateDriverRow struct {
	DriverID    int32
	Name        string
	Email       pgtype.Text
	PhoneNumber pgtype.Text
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (CreateDriverRow, error) {
	row := q.db.QueryRow(ctx, createDriver, arg.Name, arg.Email, arg.PhoneNumber)
	var i CreateDriverRow
	err := row.Scan(
		&i.DriverID,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO Orders (account_id, quotation_id) 
VALUES ($1, $2) 
RETURNING order_id, account_id, quotation_id, order_status, created_at
`

type CreateOrderParams struct {
	AccountID   int32
	QuotationID int32
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.AccountID, arg.QuotationID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.AccountID,
		&i.QuotationID,
		&i.OrderStatus,
		&i.CreatedAt,
	)
	return i, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO OrderItems (order_id, inventory_item_id, quantity) 
VALUES ($1, $2, $3) 
RETURNING order_item_id, order_id, inventory_item_id, quantity
`

type CreateOrderItemParams struct {
	OrderID         int32
	InventoryItemID int32
	Quantity        pgtype.Int4
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (Orderitem, error) {
	row := q.db.QueryRow(ctx, createOrderItem, arg.OrderID, arg.InventoryItemID, arg.Quantity)
	var i Orderitem
	err := row.Scan(
		&i.OrderItemID,
		&i.OrderID,
		&i.InventoryItemID,
		&i.Quantity,
	)
	return i, err
}

const createQuotation = `-- name: CreateQuotation :one
INSERT INTO QuotationRequest (total_cost, delivery_frequency, destination, special_handling, insurance, include_insurance, is_refused, cart_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING quotation_id, total_cost, delivery_frequency, destination, special_handling, insurance, include_insurance, is_refused, cart_id
`

type CreateQuotationParams struct {
	TotalCost         pgtype.Numeric
	DeliveryFrequency pgtype.Text
	Destination       pgtype.Text
	SpecialHandling   pgtype.Text
	Insurance         pgtype.Numeric
	IncludeInsurance  pgtype.Bool
	IsRefused         pgtype.Bool
	CartID            pgtype.Int4
}

func (q *Queries) CreateQuotation(ctx context.Context, arg CreateQuotationParams) (Quotationrequest, error) {
	row := q.db.QueryRow(ctx, createQuotation,
		arg.TotalCost,
		arg.DeliveryFrequency,
		arg.Destination,
		arg.SpecialHandling,
		arg.Insurance,
		arg.IncludeInsurance,
		arg.IsRefused,
		arg.CartID,
	)
	var i Quotationrequest
	err := row.Scan(
		&i.QuotationID,
		&i.TotalCost,
		&i.DeliveryFrequency,
		&i.Destination,
		&i.SpecialHandling,
		&i.Insurance,
		&i.IncludeInsurance,
		&i.IsRefused,
		&i.CartID,
	)
	return i, err
}

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
INSERT INTO Accounts (name, last_name, password, phone_number, email, address, authLevel) 
VALUES ($1,$2,$3,$4,$5,$6, "Client") RETURNING account_id, name, last_name, password, phone_number, email, address, authlevel
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
UPDATE Inventory SET stock_quantity = stock_quantity - $1 
WHERE inventory_id = $2 RETURNING inventory_id, pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc
`

type DecrementInventoryStockParams struct {
	StockQuantity int32
	InventoryID   int32
}

func (q *Queries) DecrementInventoryStock(ctx context.Context, arg DecrementInventoryStockParams) (Inventory, error) {
	row := q.db.QueryRow(ctx, decrementInventoryStock, arg.StockQuantity, arg.InventoryID)
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

const deleteAllCartItems = `-- name: DeleteAllCartItems :one
DELETE FROM ShoppingCartItems 
WHERE cart_id = $1 RETURNING shopping_cart_item_id, cart_id, inventory_item_id
`

func (q *Queries) DeleteAllCartItems(ctx context.Context, cartID int32) (Shoppingcartitem, error) {
	row := q.db.QueryRow(ctx, deleteAllCartItems, cartID)
	var i Shoppingcartitem
	err := row.Scan(&i.ShoppingCartItemID, &i.CartID, &i.InventoryItemID)
	return i, err
}

const deleteCart = `-- name: DeleteCart :one
DELETE FROM ShoppingCart 
WHERE cart_id = $1 RETURNING cart_id, account_id
`

func (q *Queries) DeleteCart(ctx context.Context, cartID int32) (Shoppingcart, error) {
	row := q.db.QueryRow(ctx, deleteCart, cartID)
	var i Shoppingcart
	err := row.Scan(&i.CartID, &i.AccountID)
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

const deleteQuotation = `-- name: DeleteQuotation :exec
DELETE FROM QuotationRequest
WHERE quotation_id = $1
`

func (q *Queries) DeleteQuotation(ctx context.Context, quotationID int32) error {
	_, err := q.db.Exec(ctx, deleteQuotation, quotationID)
	return err
}

const freeAllCartItems = `-- name: FreeAllCartItems :one
UPDATE InventoryItems AS ii
SET reserved = 0
FROM ShoppingCartItems AS sci
WHERE ii.inventory_item_id = sci.inventory_item_id
AND sci.cart_id = $1 RETURNING shopping_cart_item_id, cart_id, sci.inventory_item_id, ii.inventory_item_id, inventory_id, reserved
`

type FreeAllCartItemsRow struct {
	ShoppingCartItemID int32
	CartID             int32
	InventoryItemID    int32
	InventoryItemID_2  int32
	InventoryID        int32
	Reserved           pgtype.Bits
}

func (q *Queries) FreeAllCartItems(ctx context.Context, cartID int32) (FreeAllCartItemsRow, error) {
	row := q.db.QueryRow(ctx, freeAllCartItems, cartID)
	var i FreeAllCartItemsRow
	err := row.Scan(
		&i.ShoppingCartItemID,
		&i.CartID,
		&i.InventoryItemID,
		&i.InventoryItemID_2,
		&i.InventoryID,
		&i.Reserved,
	)
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

const getAllClientOrders = `-- name: GetAllClientOrders :many
SELECT order_id, account_id, quotation_id, order_status, created_at FROM Orders 
WHERE account_id = $1
`

func (q *Queries) GetAllClientOrders(ctx context.Context, accountID int32) ([]Order, error) {
	rows, err := q.db.Query(ctx, getAllClientOrders, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.AccountID,
			&i.QuotationID,
			&i.OrderStatus,
			&i.CreatedAt,
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

const getAvailableOrders = `-- name: GetAvailableOrders :many
SELECT 
    order_id,
    account_id,
    quotation_id,
    order_status,
    created_at
FROM 
    Orders o
WHERE 
    NOT EXISTS (
        SELECT 1
        FROM driver_orders d
        WHERE d.order_id = o.order_id
    )
    AND order_status = 'Created'
`

func (q *Queries) GetAvailableOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.Query(ctx, getAvailableOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.AccountID,
			&i.QuotationID,
			&i.OrderStatus,
			&i.CreatedAt,
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

const getDriverByEmail = `-- name: GetDriverByEmail :one

SELECT 
    driver_id, 
    name, 
    email, 
    phone_number,
    password
FROM 
    drivers
WHERE 
    email = $1
LIMIT 1
`

type GetDriverByEmailRow struct {
	DriverID    int32
	Name        string
	Email       pgtype.Text
	PhoneNumber pgtype.Text
	Password    string
}

// Returning the order_id and order_status
func (q *Queries) GetDriverByEmail(ctx context.Context, email pgtype.Text) (GetDriverByEmailRow, error) {
	row := q.db.QueryRow(ctx, getDriverByEmail, email)
	var i GetDriverByEmailRow
	err := row.Scan(
		&i.DriverID,
		&i.Name,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
	)
	return i, err
}

const getEmailFromOrder = `-- name: GetEmailFromOrder :one
SELECT a.email 
FROM Orders o
JOIN Accounts a ON o.account_id = a.account_id
WHERE o.order_id = $1
LIMIT 1
`

func (q *Queries) GetEmailFromOrder(ctx context.Context, orderID int32) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getEmailFromOrder, orderID)
	var email pgtype.Text
	err := row.Scan(&email)
	return email, err
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

const getOrderStatus = `-- name: GetOrderStatus :one
SELECT order_status FROM Orders WHERE order_id = $1
`

func (q *Queries) GetOrderStatus(ctx context.Context, orderID int32) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getOrderStatus, orderID)
	var order_status pgtype.Text
	err := row.Scan(&order_status)
	return order_status, err
}

const getOrdersByDriver = `-- name: GetOrdersByDriver :many
SELECT 
    o.order_id,
    o.account_id,
    o.quotation_id,
    o.order_status,
    o.created_at
FROM 
    Orders o
JOIN 
    Driver_orders d ON o.order_id = d.order_id
WHERE 
    d.driver_id = $1
`

func (q *Queries) GetOrdersByDriver(ctx context.Context, driverID int32) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersByDriver, driverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.AccountID,
			&i.QuotationID,
			&i.OrderStatus,
			&i.CreatedAt,
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

const getQuotationByID = `-- name: GetQuotationByID :one
SELECT quotation_id, total_cost, delivery_frequency, destination, special_handling, insurance, include_insurance, is_refused, cart_id FROM QuotationRequest
WHERE quotation_id = $1 LIMIT 1
`

func (q *Queries) GetQuotationByID(ctx context.Context, quotationID int32) (Quotationrequest, error) {
	row := q.db.QueryRow(ctx, getQuotationByID, quotationID)
	var i Quotationrequest
	err := row.Scan(
		&i.QuotationID,
		&i.TotalCost,
		&i.DeliveryFrequency,
		&i.Destination,
		&i.SpecialHandling,
		&i.Insurance,
		&i.IncludeInsurance,
		&i.IsRefused,
		&i.CartID,
	)
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

const getShoppingCartItemsWithPrice = `-- name: GetShoppingCartItemsWithPrice :many
SELECT 
    sc.cart_id, 
    sci.shopping_cart_item_id, 
    ii.inventory_item_id, 
    i.unit_price
FROM 
    ShoppingCart AS sc
JOIN 
    ShoppingCartItems AS sci ON sc.cart_id = sci.cart_id
JOIN 
    InventoryItems AS ii ON sci.inventory_item_id = ii.inventory_item_id
JOIN 
    Inventory AS i ON ii.inventory_id = i.inventory_id
WHERE 
    sc.cart_id = $1
`

type GetShoppingCartItemsWithPriceRow struct {
	CartID             int32
	ShoppingCartItemID int32
	InventoryItemID    int32
	UnitPrice          pgtype.Numeric
}

func (q *Queries) GetShoppingCartItemsWithPrice(ctx context.Context, cartID int32) ([]GetShoppingCartItemsWithPriceRow, error) {
	rows, err := q.db.Query(ctx, getShoppingCartItemsWithPrice, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetShoppingCartItemsWithPriceRow
	for rows.Next() {
		var i GetShoppingCartItemsWithPriceRow
		if err := rows.Scan(
			&i.CartID,
			&i.ShoppingCartItemID,
			&i.InventoryItemID,
			&i.UnitPrice,
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
UPDATE Inventory SET stock_quantity = stock_quantity + $1 
WHERE inventory_id = $2 RETURNING inventory_id, pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc
`

type IncrementInventoryStockParams struct {
	StockQuantity int32
	InventoryID   int32
}

func (q *Queries) IncrementInventoryStock(ctx context.Context, arg IncrementInventoryStockParams) (Inventory, error) {
	row := q.db.QueryRow(ctx, incrementInventoryStock, arg.StockQuantity, arg.InventoryID)
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

const markOrderAsDelivered = `-- name: MarkOrderAsDelivered :one


UPDATE Orders
SET 
    order_status = 'Delivered'
WHERE 
    order_id = $1
    AND order_status = 'Out for Delivery'  -- Ensure the order is out for delivery before marking as delivered
RETURNING order_id, order_status
`

type MarkOrderAsDeliveredRow struct {
	OrderID     int32
	OrderStatus pgtype.Text
}

// Driver ID
func (q *Queries) MarkOrderAsDelivered(ctx context.Context, orderID int32) (MarkOrderAsDeliveredRow, error) {
	row := q.db.QueryRow(ctx, markOrderAsDelivered, orderID)
	var i MarkOrderAsDeliveredRow
	err := row.Scan(&i.OrderID, &i.OrderStatus)
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
