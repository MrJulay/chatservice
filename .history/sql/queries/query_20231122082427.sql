-- name: CreateChat :exec
INSERT INTO chats
    (id, user_id, initial_message_id, status, token_usage, model, model_max_tokens, temperature, top_p, n, stop, max_tokens, presence_penalty, frequence_penalty, created_at, 

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;