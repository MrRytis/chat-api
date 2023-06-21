-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE groups
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    name      VARCHAR(255) NOT NULL,
    uuid      VARCHAR(255) NOT NULL UNIQUE,
    admin_id    INTEGER,
    CONSTRAINT fk_admin FOREIGN KEY(admin_id) REFERENCES users(id)
);

CREATE TABLE group_users
(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE groups;
DROP TABLE group_users;

