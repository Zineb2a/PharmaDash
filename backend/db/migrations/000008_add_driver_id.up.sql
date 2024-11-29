ALTER TABLE Orders ADD COLUMN driver_id INT NULL;
ALTER TABLE Orders ADD CONSTRAINT fk_driver FOREIGN KEY (driver_id) REFERENCES Accounts(account_id);