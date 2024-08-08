-- +goose Up
-- +goose StatementBegin
-- Use wal mode
CREATE TABLE IF NOT EXISTS product (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  price REAL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_image (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  product_id INTEGER NOT NULL,
  url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TRIGGER update_product_updated_at
AFTER UPDATE ON product
FOR EACH ROW
BEGIN
   UPDATE product SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;


CREATE TRIGGER update_product_image_updated_at
AFTER UPDATE ON product_image
FOR EACH ROW
BEGIN
   UPDATE product_image SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- Drop tables if they exist
DROP TABLE IF EXISTS product_image;
DROP TABLE IF EXISTS product;
-- +goose StatementEnd
