-- name: CreateChat :exec
INSERT INTO chats
    (id, user_id, )

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;