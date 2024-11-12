-- name: GetUserByEmail :one
SELECT * FROM Accounts
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO Accounts (name, last_name, password, phone_number, email, address) 
VALUES ($1,$2,$3,$4,$5,$6) RETURNING *;

-- name: GetShoppingCartByClientID :one
SELECT * FROM ShoppingCart
WHERE account_id = $1 LIMIT 1;

-- name: CreateShoppingCart :one
INSERT INTO ShoppingCart (account_id) 
VALUES ($1) RETURNING *;

-- name: GetAllShoppingCartItems :many
SELECT *
FROM ShoppingCartItems AS sci
INNER JOIN InventoryItems AS ii ON sci.inventory_item_id = ii.inventory_item_id
INNER JOIN Inventory AS i ON ii.inventory_id = i.inventory_id
WHERE sci.cart_id = $1;

-- name: GetAvailableInventoryItem :one
SELECT * FROM InventoryItems
WHERE inventory_item_id = $1
AND reserved = 0 LIMIT 1;

-- name: GetInventoryItemByID :one
SELECT * FROM InventoryItems
WHERE inventory_item_id = $1 LIMIT 1;

-- name: ReserveItem :one
UPDATE InventoryItems SET reserved = 1
WHERE inventory_item_id = $1 RETURNING *;

-- name: FreeItem :one
UPDATE InventoryItems SET reserved = 0
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
SET reserved = 0
FROM ShoppingCartItems AS sci
WHERE ii.inventory_item_id = sci.inventory_item_id
AND sci.cart_id = $1 RETURNING *;