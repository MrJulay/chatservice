-- name: CreateChat :exec


-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;