-- name: GetPiggy :one
SELECT * FROM "piggy"
WHERE address = $1;

-- name: GetPiggyFromProfileAddress :one
SELECT * FROM "piggy"
WHERE profile_address = $1;

-- name: GetPiggyFromName :one
SELECT * FROM "piggy"
WHERE name = $1;

-- name: GetPiggies :many
SELECT * FROM "piggy"
WHERE profile_address = $1;

-- name: CreatePiggy :one
INSERT INTO "piggy" (
  "address", "from_address", "profile_address", "created_at", "added_at", "unlocks_at", "name"
) VALUES ($1, $2, $3, $4, now(), $5, $6) RETURNING *;


-- name: UpdatePiggy :one
UPDATE "piggy"
SET 
  name = $1
WHERE address = $2 RETURNING *;