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
