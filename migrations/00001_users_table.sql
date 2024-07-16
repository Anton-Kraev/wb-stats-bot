-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    chat_id BIGINT PRIMARY KEY,
    token TEXT DEFAULT NULL,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
