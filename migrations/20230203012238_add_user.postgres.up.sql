CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Etc/UTC";

CREATE SCHEMA IF NOT EXISTS management;


CREATE TABLE management.roles(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO management.roles(name) VALUES
  ('admin'),
  ('accountant');

CREATE TABLE management.users(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY,
  role VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (role) REFERENCES management.roles (name)
);

INSERT INTO management.users(name, password, role) VALUES
  ('riyan', '$2a$10$k0AkHQnVAvofMRC7F3Qi2eHg20RiKuLeZW8x3hAun.rMGZcH6XEvK', 'admin'),
  ('febri', '$2a$10$k0AkHQnVAvofMRC7F3Qi2eHg20RiKuLeZW8x3hAun.rMGZcH6XEvK', 'accountant');

CREATE TABLE management.refresh_token(
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR NOT NULL,
  valid_until TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES management.users (id)
);