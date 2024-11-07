CREATE TABLE IF NOT EXISTS Accounts (
  account_id SERIAL PRIMARY KEY,
  name varchar(50) NOT NULL,
  last_name varchar(50) NOT NULL,
  password varchar(100) NOT NULL,
  phone_number varchar(100) NULL,
  email varchar(50) UNIQUE NULL,
  address varchar(50) NULL
);