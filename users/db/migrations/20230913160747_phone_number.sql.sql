-- migrate:up
ALTER TABLE users ADD COLUMN phone_number VARCHAR(20);

-- migrate:down
ALTER TABLE users DROP COLUMN phone_number;
