-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds 
ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementEnd
