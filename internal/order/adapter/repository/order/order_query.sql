-- name: InsertOrder :one
INSERT INTO
    orders (
        user_id,
        total_price,
        created_at
    )
VALUES ($1, $2, $3) RETURNING id;

-- name: InsertOrderAddress :one
INSERT INTO
    order_addresses (
        order_id,
        street,
        city,
        zip_code,
        latitude,
        longitude
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
