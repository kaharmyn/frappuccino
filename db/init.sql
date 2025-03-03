-- Drop existing tables if they exist (for development purposes)
DROP TABLE IF EXISTS order_items, orders, menu_price_history, menu_items, inventory_transactions, inventory, ingredients, menu_item_ingredients, customers;

-- Customers table (optional for future features)
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE,
    phone TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Menu Items table (should be created before order_items)
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    current_price DECIMAL(10,2) NOT NULL CHECK (current_price >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE SET NULL,
    status TEXT CHECK (status IN ('pending', 'processing', 'completed', 'canceled')) DEFAULT 'pending',
    total_price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Order Items table (menu_items now exists)
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_order DECIMAL(10,2) NOT NULL, -- Captures the price at the time of ordering
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Menu Price History table
CREATE TABLE menu_price_history (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    old_price DECIMAL(10,2) NOT NULL,
    new_price DECIMAL(10,2) NOT NULL,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory table
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    quantity INT NOT NULL CHECK (quantity >= 0),
    unit TEXT NOT NULL, -- e.g., grams, liters, pieces
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory Transactions table
CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    inventory_id INT REFERENCES inventory(id) ON DELETE CASCADE,
    change_type TEXT CHECK (change_type IN ('restock', 'usage', 'wastage')) NOT NULL,
    amount INT NOT NULL CHECK (amount <> 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Ingredients table
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Menu Item Ingredients table
CREATE TABLE menu_item_ingredients (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    ingredient_id INT REFERENCES ingredients(id) ON DELETE CASCADE,
    quantity_required INT NOT NULL CHECK (quantity_required > 0),
    unit TEXT NOT NULL -- e.g., grams, ml
);

-- Indexes for optimization
CREATE INDEX idx_orders_status ON orders (status);
CREATE INDEX idx_orders_customer ON orders (customer_id);
CREATE INDEX idx_orders_created_at ON orders USING BRIN (created_at);

CREATE INDEX idx_order_items_order_id ON order_items (order_id);
CREATE INDEX idx_order_items_menu_item_id ON order_items (menu_item_id);

CREATE INDEX idx_menu_items_name ON menu_items USING GIN (to_tsvector('english', name));
CREATE INDEX idx_menu_price_history_menu_item ON menu_price_history (menu_item_id);
CREATE INDEX idx_menu_price_history_changed_at ON menu_price_history USING BRIN (changed_at);

CREATE INDEX idx_inventory_name ON inventory (name);
CREATE INDEX idx_inventory_transactions_inventory_id ON inventory_transactions (inventory_id);
CREATE INDEX idx_inventory_transactions_created_at ON inventory_transactions USING BRIN (created_at);

CREATE INDEX idx_customers_email ON customers (LOWER(email));
CREATE INDEX idx_customers_phone ON customers (phone);

CREATE INDEX idx_menu_item_ingredients_menu_item_id ON menu_item_ingredients (menu_item_id);
CREATE INDEX idx_menu_item_ingredients_ingredient_id ON menu_item_ingredients (ingredient_id);
