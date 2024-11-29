-- name: GetUserByEmail :one
SELECT * FROM Accounts
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO Accounts (name, last_name, password, phone_number, email, address, authlevel) 
VALUES ($1,$2,$3,$4,$5,$6, 'Client') RETURNING *;

-- name: GetShoppingCartByClientID :one
SELECT * FROM ShoppingCart
WHERE account_id = $1 LIMIT 1;

-- name: CreateShoppingCart :one
INSERT INTO ShoppingCart (account_id) 
VALUES ($1) RETURNING *;

-- name: GetAllShoppingCartItems :many
SELECT i.unit_price, ii.inventory_item_id
FROM ShoppingCartItems AS sci
INNER JOIN InventoryItems AS ii ON sci.inventory_item_id = ii.inventory_item_id
INNER JOIN Inventory AS i ON ii.inventory_id = i.inventory_id
WHERE sci.cart_id = $1;

-- name: GetAvailableInventoryItem :one
SELECT * FROM InventoryItems
WHERE inventory_id = $1
AND reserved = false LIMIT 1;

-- name: GetInventoryItemByID :one
SELECT * FROM InventoryItems
WHERE inventory_item_id = $1 LIMIT 1;

-- name: ReserveItem :one
UPDATE InventoryItems SET reserved = true
WHERE inventory_item_id = $1 RETURNING *;

-- name: FreeItem :one
UPDATE InventoryItems SET reserved = false
WHERE inventory_item_id = $1 RETURNING *;

-- name: DecrementInventoryStock :one
UPDATE Inventory SET stock_quantity = stock_quantity - $1 
WHERE inventory_id = $2 RETURNING *;

-- name: IncrementInventoryStock :one
UPDATE Inventory SET stock_quantity = stock_quantity + $1 
WHERE inventory_id = $2 RETURNING *;

-- name: CreateShoppingCartItem :one
INSERT INTO ShoppingCartItems (cart_id, inventory_item_id) 
VALUES ($1,$2) RETURNING *;

-- name: DeleteCartItem :one
DELETE FROM ShoppingCartItems 
WHERE shopping_cart_item_id = $1 RETURNING *;

-- name: DeleteAllCartItems :one
DELETE FROM ShoppingCartItems 
WHERE cart_id = $1 RETURNING *;

-- name: DeleteCart :one
DELETE FROM ShoppingCart 
WHERE cart_id = $1 RETURNING *;

-- name: FreeAllCartItems :one
UPDATE InventoryItems AS ii
SET reserved = false
FROM ShoppingCartItems AS sci
WHERE ii.inventory_item_id = sci.inventory_item_id
AND sci.cart_id = $1
AND EXISTS (SELECT 1 FROM ShoppingCartItems WHERE cart_id = $1)
RETURNING ii.*;

-- name: GetShoppingCartItemsWithPrice :many
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
    sc.cart_id = $1;


-- name: CreateQuotation :one
INSERT INTO QuotationRequest (total_cost, delivery_frequency, destination, special_handling, insurance, include_insurance, cart_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetQuotationByID :one
SELECT * FROM QuotationRequest
WHERE quotation_id = $1 LIMIT 1;

-- name: CreateOrder :one
INSERT INTO Orders (account_id, delivery_frequency, destination, special_handling, include_insurance, total_cost, insurance) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING order_id;


-- name: CreateOrderItem :one
INSERT INTO OrderItems (order_id, inventory_item_id, quantity) 
VALUES ($1, $2, $3) 
RETURNING order_item_id, order_id, inventory_item_id, quantity;

-- name: DeleteQuotation :exec
DELETE FROM QuotationRequest
WHERE quotation_id = $1;


-- Update the order status to 'Out for Delivery'
UPDATE Orders
SET 
    order_status = 'Out for delivery'
WHERE 
    order_id = $2
    AND order_status = 'Created';  -- Ensure the order is still pending before assigning

COMMIT;


-- name: MarkOrderAsDelivered :one
UPDATE Orders
SET 
    order_status = 'Delivered'
WHERE 
    order_id = $1
    AND order_status = 'Out for delivery'  -- Ensure the order is out for delivery before marking as delivered
RETURNING order_id, order_status;  -- Returning the order_id and order_status


-- name: GetEmailFromOrder :one
SELECT a.email 
FROM Orders o
JOIN Accounts a ON o.account_id = a.account_id
WHERE o.order_id = $1
LIMIT 1;

-- name: GetAllClientOrders :many
SELECT * FROM Orders 
WHERE account_id = $1;

-- name: GetDriverByEmail :one
SELECT * 
FROM Accounts 
WHERE email = $1 AND authLevel = 'Driver';

-- name: GetAvailableOrders :many
SELECT * 
FROM Orders 
WHERE driver_id IS NULL AND order_status = 'Created';

-- name: AssignOrderToDriver :exec
UPDATE Orders
SET driver_id = $1, order_status = 'Out for delivery' -- The account ID of the driver
WHERE order_id = $2 AND driver_id IS NULL;

-- name: GetOrdersByDriver :many
SELECT * FROM Orders WHERE driver_id = $1;

-- name: GetQuotationByCartID :one
SELECT * FROM QuotationRequest WHERE cart_id = $1;

-- name: DeleteQuotationByCartID :exec
DELETE FROM QuotationRequest WHERE cart_id = $1;

-- name: AddFeedback :one
INSERT INTO Feedback (order_id, client_id, rating, comment)
VALUES ($1, $2, $3, $4)
RETURNING feedback_id, order_id, client_id, rating, comment, created_at;

-- name: GetOrderStatus :one
SELECT order_status FROM Orders WHERE order_id = $1;

-- name: GetAllOrders :many
SELECT * FROM Orders;

-- name: GetAllInventoryItems :many
SELECT * FROM Inventory;

-- name: AddItemToInventory :one
INSERT INTO Inventory (pharmacy_id, item_name, item_description, medication_name, unit_price, stock_quantity, otc)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING inventory_id;

-- name: AddLineItemToInventory :one
INSERT INTO InventoryItems (inventory_id, reserved)
VALUES ($1, $2)
RETURNING *;

