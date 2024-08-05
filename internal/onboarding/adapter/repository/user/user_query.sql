-- name: InsertUser :one
INSERT INTO
    users (
        name,
        phone_number,
        role,
        created_at
    )
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: FindUserByID :one
SELECT
    id,
    name,
    phone_number,
    role,
    created_at
FROM users
WHERE
    id = $1;

-- name: GetUserAddresses :many
SELECT
    id,
    user_id,
    street,
    city,
    zip_code,
    latitude,
    longitude
FROM addresses
WHERE
    user_id = $1;

-- name: InsertAddress :one
INSERT INTO
    addresses (
        user_id,
        street,
        city,
        zip_code,
        latitude,
        longitude
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
