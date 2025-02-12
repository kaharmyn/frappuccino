**Data Models**

Define data models for each entity, ensuring they are serializable to JSON and include all necessary fields.

- **Order (`models/order.go`):** 

```go
type Order struct {
    ID           string       `json:"order_id"`
    CustomerName string       `json:"customer_name"`
    Items        []OrderItem  `json:"items"`
    Status       string       `json:"status"`
    CreatedAt    string       `json:"created_at"`
}

type OrderItem struct {
    ProductID string `json:"product_id"`
    Quantity  int    `json:"quantity"`
}
```

- **Menu Item (`models/menu_item.go`):**
```go
type MenuItem struct {
  ID          string                `json:"product_id"`
  Name        string                `json:"name"`
  Description string                `json:"description"`
  Price       float64               `json:"price"`
  Ingredients []MenuItemIngredient  `json:"ingredients"`
}

type MenuItemIngredient struct {
  IngredientID string `json:"ingredient_id"`
  Quantity     float64    `json:"quantity"`
}
```

- **Inventory Item (`models/inventory_item.go`):**
```go
type InventoryItem struct {
    IngredientID string `json:"ingredient_id"`
    Name         string `json:"name"`
    Quantity     float64    `json:"quantity"`
    Unit         string `json:"unit"`
}
```

### API Endpoints

Implement the following RESTful API endpoints:

- **Orders:**

  - `POST /orders`: Create a new order.
  - `GET /orders`: Retrieve all orders.
  - `GET /orders/{id}`: Retrieve a specific order by ID.
  - `PUT /orders/{id}`: Update an existing order.
  - `DELETE /orders/{id}`: Delete an order.
  - `POST /orders/{id}/close`: Close an order.

- **Menu Items:**

  - `POST /menu`: Add a new menu item.
  - `GET /menu`: Retrieve all menu items.
  - `GET /menu/{id}`: Retrieve a specific menu item.
  - `PUT /menu/{id}`: Update a menu item.
  - `DELETE /menu/{id}`: Delete a menu item.

- **Inventory:**

  - `POST /inventory`: Add a new inventory item.
  - `GET /inventory`: Retrieve all inventory items.
  - `GET /inventory/{id}`: Retrieve a specific inventory item.
  - `PUT /inventory/{id}`: Update an inventory item.
  - `DELETE /inventory/{id}`: Delete an inventory item.

- **Aggregations:**

  - `GET /reports/total-sales`: Get the total sales amount.
  - `GET /reports/popular-items`: Get a list of popular menu items.

#### Updating Inventory Upon Order Fulfillment
When an order is created and processed, the application must:

- Check Inventory Levels:
  - Before confirming an order, verify that there are sufficient quantities of all required ingredients in `inventory.json`.
  - If any ingredient is insufficient, the order should not be processed, and an appropriate error message should be returned.
- Deduct Ingredients:
  - Upon successful processing of an order, deduct the required quantities of each ingredient from `inventory.json`.

**Examples:**

- **Create Order Request:**
```http 
POST /orders
Content-Type: application/json

{
  "customer_name": "John Doe",
  "items": [
    {
      "product_id": "espresso",
      "quantity": 2
    },
    {
      "product_id": "croissant",
      "quantity": 1
    }
  ]
}
```

- **Total Sales Aggregation Response:**
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "total_sales": 1500.50
}
```

### Data Storage with JSON Files
As part of the project requirements, you must use JSON files as the data storage mechanism instead of a database. The application should store all data locally in JSON files, with each entity stored in its own separate file within the `data/` directory. 

#### Entities to Store:

- Orders: Information about customer orders.
- Menu Items: Details about the products available in the coffee shop.
- Inventory Items: Inventory of ingredients required to prepare menu items.

#### Requirements:

- **No Database Usage:** Do not use any database systems. All data must be stored and retrieved from JSON files.
- **Separate Files for Each Entity:** Each entity should have its own JSON file:
  - `orders.json` for orders.
  - `menu_items.json` for menu items.
  - `inventory.json` for inventory items.
- **Include Ingredients in Menu Items**: `menu_items.json` should include the list of ingredients required for each menu item.
- **Update Inventory on Order Fulfillment**: When an order is fulfilled, the application must update `inventory.json` by deducting the quantities of ingredients used.
- **Data Format**: Data should be stored in structured JSON formats that accurately represent the entities.

#### Examples of JSON Files:

1. `orders.json`:

```json
[
  {
    "order_id": "order123",
    "customer_name": "Alice Smith",
    "items": [
      {
        "product_id": "latte",
        "quantity": 2
      },
      {
        "product_id": "muffin",
        "quantity": 1
      }
    ],
    "status": "open",
    "created_at": "2023-10-01T09:00:00Z"
  },
  {
    "order_id": "order124",
    "customer_name": "Bob Johnson",
    "items": [
      {
        "product_id": "espresso",
        "quantity": 1
      }
    ],
    "status": "closed",
    "created_at": "2023-10-01T09:30:00Z"
  }
]
```

2. `menu_items.json`:
```json
[
  {
    "product_id": "latte",
    "name": "Caffe Latte",
    "description": "Espresso with steamed milk",
    "price": 3.50,
    "ingredients": [
      {
        "ingredient_id": "espresso_shot",
        "quantity": 1
      },
      {
        "ingredient_id": "milk",
        "quantity": 200
      }
    ]
  },
  {
    "product_id": "muffin",
    "name": "Blueberry Muffin",
    "description": "Freshly baked muffin with blueberries",
    "price": 2.00,
    "ingredients": [
      {
        "ingredient_id": "flour",
        "quantity": 100
      },
      {
        "ingredient_id": "blueberries",
        "quantity": 20
      },
      {
        "ingredient_id": "sugar",
        "quantity": 30
      }
    ]
  },
  {
    "product_id": "espresso",
    "name": "Espresso",
    "description": "Strong and bold coffee",
    "price": 2.50,
    "ingredients": [
      {
        "ingredient_id": "espresso_shot",
        "quantity": 1
      }
    ]
  }
]
```

**Note:** The ingredients field in each menu item lists the ingredients required to prepare that item. The quantity is specified in units appropriate for the ingredient (e.g., grams, milliliters).

3. `inventory.json`:

```json
[
  {
    "ingredient_id": "espresso_shot",
    "name": "Espresso Shot",
    "quantity": 500, // Number of shots
    "unit": "shots"
  },
  {
    "ingredient_id": "milk",
    "name": "Milk",
    "quantity": 5000, // In milliliters
    "unit": "ml"
  },
  {
    "ingredient_id": "flour",
    "name": "Flour",
    "quantity": 10000, // In grams
    "unit": "g"
  },
  {
    "ingredient_id": "blueberries",
    "name": "Blueberries",
    "quantity": 2000,  // In grams
    "unit": "g"
  },
  {
    "ingredient_id": "sugar",
    "name": "Sugar",
    "quantity": 5000, // In grams
    "unit": "g"
  }
]
```


### Logging
- Use Go's `log/slog` package for logging throughout the application.
- Log significant events, errors information with appropriate levels (`Info`, `Warning`, `Error`).
- Include contextual information in logs (e.g., timestamps, IDs).

**Example:**

```go
slog.Info("Order created", "orderID", newOrder.ID)
slog.Error("Failed to update inventory", err)
```

### Error Handling and Validation
- Validate input and provide meaningful error messages.
- Return appropriate HTTP status codes:
  - `200 OK` for successful GET requests.
  - `201 Created` for successful POST requests.
  - `400 Bad Request` for invalid input.
  - `404 Not Found` when resources are not found.
  - `500 Internal Server Error` for unexpected errors.
- Ensure error responses include a message explaining the error.

**Examples:**

```http 
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": "Invalid product ID in order items."
}
```

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": "Insufficient inventory for ingredient 'Milk'. Required: 200ml, Available: 150ml."
}

```