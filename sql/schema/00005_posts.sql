-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title VARCHAR(255) NOT NULL,
  url TEXT UNIQUE NOT NULL,
  description TEXT NOT NULL,
  published_at TIMESTAMP NOT NULL,
  feed_id UUID NOT NULL,
  FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd