-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    uuid       varchar(36) NOT NULL UNIQUE,
    user_id    INTEGER     NOT NULL,
    group_id   INTEGER     NOT NULL,
    content    TEXT        NOT NULL,
    is_edited  BOOLEAN     NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE messages;
