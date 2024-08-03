-- +migrate Up
CREATE TABLE categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(255) NOT NULL,
	modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_by VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE categories;
