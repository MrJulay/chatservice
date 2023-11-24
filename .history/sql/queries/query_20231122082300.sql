-- name: CreateChat :exec
INSERT INTO chats
    

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;