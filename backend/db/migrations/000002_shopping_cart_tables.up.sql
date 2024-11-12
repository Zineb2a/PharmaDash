ALTER TABLE Accounts
ADD authLevel varchar(50) NOT NULL;

CREATE TABLE IF NOT EXISTS Pharmacies (
  pharmacy_id SERIAL PRIMARY KEY,
  company_name varchar(50) NOT NULL,
  location varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Inventory (
  inventory_id SERIAL PRIMARY KEY,
  pharmacy_id int NOT NULL,
  item_name varchar(100) NOT NULL,
  item_description varchar(80) NOT NULL,
  medication_name varchar(50) NOT NULL,
  unit_price money NOT NULL,
  stock_quantity int NOT NULL,
  otc bit NOT NULL,
  FOREIGN KEY (pharmacy_id) REFERENCES Pharmacies(pharmacy_id)
);

CREATE TABLE IF NOT EXISTS InventoryItems (
  inventory_item_id SERIAL PRIMARY KEY,
  inventory_id int NOT NULL,
  reserved bit NOT NULL,
  FOREIGN KEY (inventory_id) REFERENCES Inventory(inventory_id)
);

CREATE TABLE IF NOT EXISTS Prescriptions (
  prescription_id SERIAL PRIMARY KEY,
  inventory_id int NOT NULL,
  client_email varchar(100) NOT NULL,
  doctor_name varchar(80) NOT NULL,
  medication_name varchar(50) NOT NULL,
  dose_quantity int NOT NULL,
  issued_date date NOT NULL,
  expiry_date date NOT NULL,
  FOREIGN KEY (inventory_id) REFERENCES Inventory(inventory_id)
);

CREATE TABLE IF NOT EXISTS ShoppingCart (
  cart_id SERIAL PRIMARY KEY,
  account_id int NOT NULL,
  FOREIGN KEY (account_id) REFERENCES Accounts(account_id)
);

CREATE TABLE IF NOT EXISTS ShoppingCartItems (
  shopping_cart_item_id SERIAL PRIMARY KEY,
  cart_id int NOT NULL,
  inventory_item_id int NOT NULL,
  FOREIGN KEY (cart_id) REFERENCES ShoppingCart(cart_id),
  FOREIGN KEY (inventory_item_id) REFERENCES InventoryItems(inventory_item_id)
);