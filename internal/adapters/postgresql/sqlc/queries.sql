-- name: ListProducts :many
SELECT * 
FROM products;

-- name: FindProductByID :one
SELECT *
FROM products
WHERE id = $1;

-- name: PlaceOrder :exec