CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users (
    id string PRIMARY key,
    name string
, phone_number VARCHAR(20));
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20221202132622'),
  ('20230913160747');
