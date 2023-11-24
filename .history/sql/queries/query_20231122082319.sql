-- name: CreateChat :exec
INSERT INTO chats
    (id, user_id, initial_message_id, status, token_usage, model, )

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;