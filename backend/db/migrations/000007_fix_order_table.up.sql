ALTER TABLE Orders DROP COLUMN quotation_id;

ALTER TABLE Orders ADD COLUMN destination TEXT;
ALTER TABLE Orders ADD COLUMN special_handling TEXT;
ALTER TABLE Orders ADD COLUMN delivery_frequency TEXT;
ALTER TABLE Orders ADD COLUMN include_insurance BOOLEAN;
ALTER TABLE Orders ADD COLUMN total_cost FLOAT;
ALTER TABLE Orders ADD COLUMN insurance FLOAT;


    --total_cost NUMERIC,
    -- delivery_frequency TEXT,
    -- destination TEXT,
    -- special_handling TEXT,
    --insurance NUMERIC,
    -- include_insurance BOOLEAN,