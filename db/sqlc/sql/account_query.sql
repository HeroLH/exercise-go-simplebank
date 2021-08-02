-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1 LIMIT 1;

-- 等待更新后再获得值
-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;
-- NO KEY 告诉 Pgsql 我们不更新帐户表的键或 ID 列


-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id LIMIT $2
OFFSET $3;