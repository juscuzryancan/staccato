-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tournaments (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  rawData JSONB 
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tournaments;
-- +goose StatementEnd


