-- Create ENUM types
CREATE TYPE order_status AS ENUM ('open', 'closed', 'cancelled');
CREATE TYPE unit_type AS ENUM ('g', 'ml', 'shots', 'pieces');
-- CREATE TYPE payment_method AS ENUM ('cash', 'card', 'online');

-- Create tables

-- INVENTORY SERVICE
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    quantity NUMERIC(10, 2) NOT NULL,
    unit unit_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
); -- DONE

CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    ingredient_id INT REFERENCES inventory(id) ON DELETE CASCADE,
    quantity NUMERIC(10, 2) NOT NULL,
    transaction_type TEXT NOT NULL CHECK (transaction_type IN ('add', 'deduct')),
    transaction_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- MENU SERVICE
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    categories TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE menu_item_ingredients (
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    ingredient_id INT REFERENCES inventory(id) ON DELETE CASCADE,
    quantity NUMERIC(10, 2) NOT NULL,
    PRIMARY KEY (menu_item_id, ingredient_id)
); -- check this out with PK queries

-- ORDER SERVICE
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    status order_status NOT NULL DEFAULT 'open',
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    price_at_order NUMERIC(10, 2) NOT NULL,
    PRIMARY KEY (order_id, menu_item_id)
);


--
CREATE TABLE order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    status order_status NOT NULL,
    changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
--

CREATE TABLE price_history (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    price NUMERIC(10, 2) NOT NULL,
    changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


-- Create indexes
CREATE INDEX idx_menu_items_name ON menu_items(name);
CREATE INDEX idx_orders_customer_name ON orders(customer_name);
CREATE INDEX idx_inventory_name ON inventory(name);
CREATE INDEX idx_order_items_order_id ON order_items(order_id);

-- Insert mock data
-- Menu items
INSERT INTO menu_items (name, description, price, categories)
VALUES
    ('Caffe Latte', 'Espresso with steamed milk', 3.50, ARRAY['coffee', 'hot']),
    ('Blueberry Muffin', 'Freshly baked muffin with blueberries', 2.00, ARRAY['bakery', 'snack']),
   