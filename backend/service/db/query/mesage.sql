-- name: AddMessage :one
INSERT INTO "message" (
  "transaction_hash", "address", "token", "amount", "fee", "content", "nick"
) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetMessageByHash :one
SELECT * FROM "message"
WHERE "transaction_hash" = $1;

-- name: GetMessagesByAddress :many
SELECT * FROM "message"
WHERE "address" = $1;