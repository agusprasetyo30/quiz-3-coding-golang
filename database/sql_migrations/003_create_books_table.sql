-- +migrate Up
CREATE TABLE books (
	id SERIAL PRIMARY KEY,
	category_id INTEGER NOT NULL,
	description TEXT NULL,
	image_url VARCHAR(255) NOT NULL,
	release_year INTEGER NULL,
	price INTEGER NOT NULL,
	total_page INTEGER NOT NULL,
	thickness VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(255) NULL,
	modified_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	modified_by VARCHAR(255) NULL,
	CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id)
);

-- +migrate Down
DROP TABLE books;
