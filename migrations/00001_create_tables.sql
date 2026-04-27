-- +goose Up
CREATE TABLE orders (
  id TEXT PRIMARY KEY,
  user_id TEXT,
  amount DOUBLE PRECISION,
  created_at TIMESTAMP
);

CREATE TABLE outbox (
  id SERIAL PRIMARY KEY,
  topic TEXT,
  key TEXT,
  payload JSONB,
  processed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE outbox;
DROP TABLE orders;