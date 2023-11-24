-- name: CreateChat :exec
INSERT INTO chats
    (id, user_id, initial_message_id, status, token_usage, model, model_max_tokens, temperature, top_n, n)

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;