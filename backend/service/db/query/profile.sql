-- name: GetProfile :one
SELECT * FROM "profile"
WHERE address = $1;

-- name: CreateProfile :one
INSERT INTO "profile" (
  "address"
) VALUES ($1) RETURNING *;


-- name: CheckUserRegistration :one
SELECT EXISTS (
    SELECT 1
    FROM "profile"
    WHERE address = $1
) AS is_registered;