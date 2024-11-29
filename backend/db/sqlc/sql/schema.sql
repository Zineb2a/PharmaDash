CREATE TABLE IF NOT EXISTS Accounts (
  account_id SERIAL PRIMARY KEY,
  name varchar(50) NOT NULL,
  last_name varchar(50) NOT NULL,
  password varchar(100) NOT NULL,
  phone_number varchar(100) NULL,
  email varchar(50) UNIQUE NULL,
  address varchar(50) NULL,
  authLevel varchar(50) NOT NULL
);

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
  unit_price FLOAT NOT NULL,
  stock_quantity int NOT NULL,
  otc BOOLEAN NOT NULL,
  FOREIGN KEY (pharmacy_id) REFERENCES Pharmacies(pharmacy_id)
);

CREATE TABLE IF NOT EXISTS InventoryItems (
  inventory_item_id SERIAL PRIMARY KEY,
  inventory_id int NOT NULL,
  reserved BOOLEAN NOT NULL,
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

CREATE TABLE IF NOT EXISTS QuotationRequest (
    quotation_id SERIAL PRIMARY KEY,
    total_cost FLOAT,
    delivery_frequency TEXT,
    destination TEXT,
    special_handling TEXT,
    insurance FLOAT,
    include_insurance BOOLEAN,
    cart_id INTEGER,
    FOREIGN KEY (cart_id) REFERENCES ShoppingCart(cart_id)
);


CREATE TABLE IF NOT EXISTS Orders (
    order_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    order_status VARCHAR(50) DEFAULT 'Created',
    created_at TIMESTAMP DEFAULT NOW(),
    delivery_frequency TEXT,
    destination TEXT,
    special_handling TEXT,
    include_insurance BOOLEAN,
    total_cost FLOAT,
    insurance FLOAT,
    driver_id INT NULL, 
    FOREIGN KEY (account_id) REFERENCES Accounts(account_id),
    FOREIGN KEY (driver_id) REFERENCES Accounts(account_id)
);

--ALTER TABLE Orders
--ADD COLUMN driver_id INT NULL, -- The account ID of the driver assigned to the order
--ADD CONSTRAINT fk_driver FOREIGN KEY (driver_id) REFERENCES Accounts(account_id);

--ALTER TABLE Orders
--    ALTER COLUMN driver_id SET DATA TYPE UUID USING driver_id::UUID;

CREATE TABLE IF NOT EXISTS OrderItems (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    inventory_item_id INT NOT NULL,
    quantity INT DEFAULT 1,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (inventory_item_id) REFERENCES InventoryItems(inventory_item_id)
);

CREATE TABLE IF NOT EXISTS Feedback (
    feedback_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    client_id INT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5), -- Rating between 1 and 5
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (client_id) REFERENCES Accounts(account_id)
);

CREATE TABLE IF NOT EXISTS Feedback (
    feedback_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    client_id INT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5), -- Rating between 1 and 5
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (client_id) REFERENCES Accounts(account_id)
);

