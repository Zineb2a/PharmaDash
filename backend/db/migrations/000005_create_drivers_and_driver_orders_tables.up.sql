-- Create the drivers table
CREATE TABLE IF NOT EXISTS Drivers (
    driver_id SERIAL PRIMARY KEY, -- Auto-incrementing primary key
    name VARCHAR(50) NOT NULL,    -- Driver's name
    phone_number VARCHAR(100),    -- Optional phone number
    email VARCHAR(50) UNIQUE      -- Optional unique email
);

-- Create the driver_orders association table
-- This links drivers to orders, allowing each driver to have multiple orders
CREATE TABLE IF NOT EXISTS Driver_orders (
    driver_id INT NOT NULL REFERENCES drivers(driver_id) ON DELETE CASCADE,
    order_id INT NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    PRIMARY KEY (driver_id, order_id) -- Prevents duplicate associations
);
