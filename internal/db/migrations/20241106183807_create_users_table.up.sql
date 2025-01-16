CREATE TABLE users (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password_hash varchar(255) NOT NULL,
  password_salt varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT users_pk PRIMARY KEY (id),
  CONSTRAINT users_email_idx UNIQUE (id)
);

CREATE TRIGGER update_users_columns BEFORE
UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_columns();
