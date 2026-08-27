# E-commerce-app-Go-Lang
A backend REST API for an E-Commerce application built with **Golang**.

This project is designed to demonstrate backend development concepts including **RESTful APIs, HTTP request handling, middleware, JSON processing, product management, and JWT-based authentication**.

The application follows a modular architecture to keep the code organized, maintainable, and easy to extend with new features.

---

##  About The Project

The **E-Commerce REST API** is a backend application developed using **Golang (Go)**. The main goal of this project is to build a structured backend system that can be used as the foundation for a complete E-Commerce platform.

The application handles client requests through RESTful API endpoints and communicates using JSON. The project is organized into separate components such as routes, handlers, middleware, database logic, and utility functions.

This architecture helps separate responsibilities and makes the application easier to understand, maintain, and scale.

The API can later be connected with frontend applications built using technologies.


##  Features

### Product Management

- Create new products
- Store product information
- Retrieve product lists
- Automatically generate product IDs
- Handle product data using Go structs
- Return product information in JSON format

### RESTful API

- REST API architecture
- HTTP request handling
- JSON request parsing
- JSON response generation
- RESTful endpoint design
- Proper HTTP method usage
- Structured request and response handling

###  Authentication

- JWT-based authentication architecture
- Secure token generation
- JWT payload handling
- User information support inside tokens
- Authentication middleware support
- Protected route architecture

###  Middleware

The project uses middleware to process HTTP requests before they reach the application handlers.

Middleware can be used for:

- Authentication
- Authorization
- Request validation
- Request logging
- Security checks
- Custom request processing

###  Modular Architecture

The project separates different responsibilities into dedicated packages.

The architecture includes:

- Routes
- Handlers
- Middleware
- Database layer
- Utility functions
- Configuration

---

##  Technology Stack

| Technology | Purpose |
|------------|---------|
| **Golang** | Backend development |
| **net/http** | HTTP server and routing |
| **JSON** | API communication |
| **JWT** | Authentication |
| **Postman** | API testing |
| **Git** | Version control |
| **GitHub** | Source code hosting |

---

# Architecture Overview

The application follows a layered architecture:

```text
Client
   │
   ▼
Routes
   │
   ▼
Middleware
   │
   ▼
Handlers
   │
   ▼
Database / Data Layer
   │
   ▼
Response
```

### Request Flow

1. The client sends an HTTP request.
2. The request reaches the application route.
3. Middleware processes the request.
4. Authentication or authorization checks can be performed.
5. The request is forwarded to the appropriate handler.
6. The handler processes the business logic.
7. The database layer stores or retrieves data.
8. A JSON response is returned to the client.

---




#  API Endpoints

##  Products

### Get All Products

Returns all available products.

**Method**

```http
GET /products
```

### Example Response

```json
[
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product Description",
    "price": 1000,
    "img_url": "https://example.com/product.jpg"
  }
]
```
![Get Products API Screenshot](screenshots/get-products.png)
- - -
##  Create Product

Creates a new product.

**Method**

```http
POST /products
```

### Example Request

```json
{
        "name": "Lemon",
        "description": "Lemon color is Green",
        "price": 60,
        "img_url": "https://images.contentstack.io/v3/assets/bltcedd8dbd5891265b/blt2a5be8abcac1a15f/667081fd5014f14c2a033ce6/types-of-cherries-on-branch.jpg"
    }
```

### Example Response

```json
{
   "id": 3
    "name": "Lemon",
   "description": "Lemon color is Green",
    "price": 60,
   "img_url": "https://images.contentstack.io/v3/assets/bltcedd8dbd5891265b/blt2a5be8abcac1a15f/667081fd5014f14c2a033ce6/types-of-         cherries-on-branch.jpg"
    
}

```
![Get Products API Screenshot](screenshots/get-products.png)
---
##  Update Product

Updates an existing product.

**Method**

```http
PUT /products/{id}
```

### Example Request

```json
{
  "name": "Updated iPhone 15",
  "description": "Updated Apple smartphone description",
  "price": 125000,
  "img_url": "https://example.com/updated-iphone.jpg"
}
```

### Example Response

```json
{
  "id": 1,
  "name": "Updated iPhone 15",
  "description": "Updated Apple smartphone description",
  "price": 125000,
  "img_url": "https://example.com/updated-iphone.jpg"
}
```
![Get Products API Screenshot](screenshots/get-products.png)

---
##  Delete Product

Deletes an existing product.

**Method**

```http
DELETE /products/{id}
```

### Example Response

```json
{
  "message": "Product deleted successfully"
}
```

![Get Products API Screenshot](screenshots/get-products.png)

---

#  JWT Authentication

The project includes JWT-based authentication functionality.

JSON Web Tokens are used to securely represent authenticated users.

The JWT payload can contain user information such as:

```json
{
  "sub": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "is_shopowner": false
}
```

The token can then be used by middleware to identify authenticated users and protect restricted API endpoints.

![Get Products API Screenshot](screenshots/get-products.png)
---
##  Authentication Testing

### JWT Token Generation


![JWT Authentication](screenshots/jwt-authentication.png)

---

#  Middleware System

The project uses a middleware manager to make middleware handling more organized.

Example architecture:

```text
HTTP Request
      │
      ▼
Middleware Manager
      │
      ▼
Authentication Middleware
      │
      ▼
Authorization Middleware
      │
      ▼
API Handler
      │
      ▼
HTTP Response
```

Middleware allows reusable request processing logic without duplicating code inside every handler.

---

# 🧪 API Testing

The APIs are tested using **Postman**.

Testing includes:

- Sending GET requests
- Sending POST requests
- Sending PUT requests
- Sending DELETE requests
- Testing JSON request bodies
- Checking API responses
- Testing HTTP status codes
- Testing authentication functionality
- Testing protected routes

---

# ▶️ Getting Started

Follow these instructions to run the project locally.

## Prerequisites

Make sure you have the following installed:

- Go 1.20 or later
- Git
- Postman (optional, for API testing)

---

## Clone the Repository

```bash
git clone YOUR_REPOSITORY_URL
```

---

## Navigate to the Project

```bash
cd E-commerce-app
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Configure Environment Variables

Create a `.env` file if required by the application.


```env
VERSION =1.0.0
SERVICE_NAME =ECOMMERCE
HTTP_PORT=3000y
```

# What I Learned From This Project

Through this project, I practiced and improved my understanding of:

- Golang programming
- Backend development
- REST API development
- HTTP methods
- HTTP request handling
- HTTP response handling
- JSON encoding and decoding
- Structs and data modeling
- Middleware architecture
- JWT authentication
- Authorization concepts
- Modular project architecture
- API testing with Postman
- Git and GitHub workflow

#  Future Improvements

The project can be extended with additional e-commerce features in the future.

Planned improvements may include:
- Shopping cart system
- Order management
- Payment integration
- Database integration
- PostgreSQL/MySQL support
- Role-based authorization
- Admin dashboard API
- Pagination
- Product search
- Product categories
- Docker containerization
- CI/CD pipeline
- Unit testing
- Deployment to cloud infrastructure
