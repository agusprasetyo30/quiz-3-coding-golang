-- 001_create_users_table.sql

-- +migrate Up
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(255) NULL,
	modified_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	modified_by VARCHAR(255) NULL
);

INSERT INTO users (username, password, created_by) VALUES
    ('admin', 'jGl25bVBBBW96Qi9Te4V37Fnqchz/Eu4qB9vKrRIqRg=', 'system'), -- Password : admin
    ('user1', 'CgQblGLKpKMbrDVn4Lbm/ZEAeH2yq0M9lvbReMq/zpA=', 'system'); -- Password : user1

-- +migrate Down
DROP TABLE users;
