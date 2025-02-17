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
-- Insert inventory items
INSERT INTO inventory (name, quantity, unit)
VALUES
    ('Espresso Beans', 5000, 'g'),
    ('Milk', 20000, 'ml'),
    ('Blueberries', 3000, 'g'),
    ('Flour', 10000, 'g'),
    ('Sugar', 5000, 'g'),
    ('Butter', 3000, 'g'),
    ('Eggs', 200, 'pieces'),
    ('Paper Cups', 500, 'pieces'),
    ('Chocolate Syrup', 1000, 'ml');

-- Insert inventory transactions
INSERT INTO inventory_transactions (ingredient_id, quantity, transaction_type)
VALUES
    (1, 1000, 'add'),
    (2, 5000, 'add'),
    (3, 1000, 'add'),
    (4, 2000, 'add'),
    (5, 1000, 'add'),
    (6, 500, 'add'),
    (7, 50, 'add'),
    (8, 100, 'add');

-- Insert additional menu items
INSERT INTO menu_items (name, description, price, categories)
VALUES
    ('Cappuccino', 'Espresso with steamed milk and foam', 3.75, ARRAY['coffee', 'hot']),
    ('Iced Coffee', 'Cold brewed coffee with ice', 2.50, ARRAY['coffee', 'cold']),
    ('Chocolate Croissant', 'Flaky croissant with chocolate filling', 3.00, ARRAY['bakery', 'snack']),
    ('Americano', 'Espresso with hot water', 2.75, ARRAY['coffee', 'hot']);

-- Insert menu item ingredients
INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity)
VALUES
    (1, 1, 18),  -- Caffe Latte: 18g Espresso Beans
    (1, 2, 150), -- Caffe Latte: 150ml Milk
    (2, 3, 50),  -- Blueberry Muffin: 50g Blueberries
    (2, 4, 100), -- Blueberry Muffin: 100g Flour
    (2, 5, 50),  -- Blueberry Muffin: 50g Sugar
    (2, 6, 30),  -- Blueberry Muffin: 30g Butter
    (2, 7, 1),   -- Blueberry Muffin: 1 Egg
    (3, 1, 18),  -- Cappuccino: 18g Espresso Beans
    (3, 2, 120), -- Cappuccino: 120ml Milk
    (4, 1, 18),  -- Iced Coffee: 18g Espresso Beans
    (4, 2, 200), -- Iced Coffee: 200ml Milk

-- Insert sample orders
INSERT INTO orders (customer_name, status, price)
VALUES
    ('John Doe', 'open', 7.50),
    ('Jane Smith', 'closed', 5.50),
    ('Alice Brown', 'cancelled', 3.75);

-- Insert order items
INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order)
VALUES
    (1, 1, 1, 3.50), -- John Doe ordered 1 Caffe Latte
    (1, 2, 2, 4.00), -- John Doe ordered 2 Blueberry Muffins
    (2, 3, 1, 3.75), -- Jane Smith ordered 1 Cappuccino
    (2, 5, 1, 3.00), -- Jane Smith ordered 1 Chocolate Croissant
    (3, 4, 1, 2.75); -- Alice Brown ordered 1 Americano

-- Insert order status history
INSERT INTO order_status_history (order_id, status)
VALUES
    (1, 'open'),
    (2, 'open'),
    (2, 'closed'),
    (3, 'open'),
    (3, 'cancelled');

-- Insert price history for menu items
INSERT INTO price_history (menu_item_id, price)
VALUES
    (1, 3.50),
    (2, 2.00),
    (3, 3.75),
    (4, 2.50),
    (5, 3.00);
