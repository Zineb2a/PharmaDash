CREATE TABLE QuotationRequest (
    quotation_id SERIAL PRIMARY KEY,
    total_cost NUMERIC,
    delivery_frequency TEXT,
    destination TEXT,
    special_handling TEXT,
    insurance NUMERIC,
    include_insurance BOOLEAN,
    is_refused BOOLEAN,
    cart_id INT
);