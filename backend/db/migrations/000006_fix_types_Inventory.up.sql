ALTER TABLE Inventory DROP COLUMN otc;
ALTER TABLE Inventory ADD COLUMN otc BOOLEAN;

ALTER TABLE Inventory DROP COLUMN unit_price;
ALTER TABLE Inventory ADD COLUMN unit_price FLOAT;

ALTER TABLE InventoryItems DROP COLUMN reserved;
ALTER TABLE InventoryItems ADD COLUMN reserved BOOLEAN;

ALTER TABLE QuotationRequest DROP COLUMN total_cost;
ALTER TABLE QuotationRequest ADD COLUMN total_cost FLOAT;

ALTER TABLE QuotationRequest DROP COLUMN insurance;
ALTER TABLE QuotationRequest ADD COLUMN insurance FLOAT;