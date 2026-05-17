-- name: CreateDeveloperProfile :one
INSERT INTO developer_profiles (name, email, phone, address, accept_terms_of_service)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, phone, address, accept_terms_of_service, created_at, updated_at;

-- name: GetDeveloperProfile :one
SELECT id, name, email, phone, address, accept_terms_of_service, created_at, updated_at
FROM developer_profiles
WHERE id = $1;

-- name: ListDeveloperProfiles :many
SELECT id, name, email, phone, address, accept_terms_of_service, created_at, updated_at
FROM developer_profiles
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateDeveloperProfile :one
UPDATE developer_profiles
SET name                    = $2,
    email                   = $3,
    phone                   = $4,
    address                 = $5,
    accept_terms_of_service = $6,
    updated_at              = NOW()
WHERE id = $1
RETURNING id, name, email, phone, address, accept_terms_of_service, created_at, updated_at;

-- name: DeleteDeveloperProfile :exec
DELETE FROM developer_profiles
WHERE id = $1;
