START TRANSACTION;
CREATE TABLE IF NOT EXISTS `chats` (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    initial_message_id TEXT NOT NULL,
    status VARCHAR(6) NOT NULL,
    token_usage SMALLINT NOT NULL,
    model VARCHAR(20) NOT NULL,
    model_max_tokens SMALLINT NOT NULL,
    temperature DECIMAL(3,2)
)