ALTER TABLE Inventory DROP COLUMN otc;
ALTER TABLE Inventory ADD COLUMN otc BOOLEAN;

ALTER TABLE Inventory DROP COLUMN unit_price;
ALTER TABLE Inventory ADD COLUMN unit_price FLOAT;

ALTER TABLE InventoryItems DROP COLUMN reserved;
ALTER TABLE InventoryItems ADD COLUMN reserved BOOLEAN;