-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE refresh_tokens
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    user_id    INTEGER,
    token      VARCHAR(255) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE refresh_tokens;
