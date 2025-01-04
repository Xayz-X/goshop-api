## Folder Structure:

```
goshop-api/
├── cmd/
│ └── main.go # Entry point of the app
├── config/
│ └── config.go # Load environment variables
├── controllers/
│ ├── auth_controller.go # Handles auth routes
│ ├── product_controller.go# Handles product routes
│ ├── cart_controller.go # Handles cart routes
│ └── order_controller.go # Handles order routes
├── models/
│ ├── user.go # User model
│ ├── product.go # Product model
│ ├── cart.go # Cart model
│ └── order.go # Order model
├── routes/
│ └── routes.go # Route definitions
├── services/
│ ├── jwt_service.go # JWT generation/validation
│ └── db_service.go # MongoDB connection
├── utils/
│ ├── hash.go # Password hashing
│ └── response.go # Utility for API responses
├── .env # Environment variables
├── go.mod # Go modules file
└── README.md # Project documentation
```

## Roues Structure:

```
Route	              Method	 Description	                Access
----------------------------------------------------------------------
/api/register	     POST	 Register a new user	          Public
/api/login	        POST	 Log in and get JWT	          Public
/api/products	     GET	    Get all products	             Public
/api/products	     POST	 Add a new product	          Admin Only
/api/products/{id}  DELETE	 Delete a product	             Admin Only
/api/cart	        POST	 Add a product to the cart	    Authenticated
/api/cart	        GET	    View items in the cart	       Authenticated
/api/orders	        POST	 Place an order	             Authenticated
/api/orders	        GET	    View user order history	    Authenticated
```

## Simplified Project: GoShop API (Lite Version)

**Core Features**

1. User Authentication:

   - User Registration.
   - User Login (JWT-based).

2. Product Management:

   - List all products (public).
   - Admin-only routes to add and delete products.

3. Cart Management:

   - Add a product to the cart.
   - View cart.

4. Order Management:
   - Place an order.
   - View user’s order history.
