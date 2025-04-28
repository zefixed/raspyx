-- +goose Up
-- +goose StatementBegin
ALTER TABLE schedule ADD is_session BOOL NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE schedule DROP COLUMN is_session;
-- +goose StatementEnd
