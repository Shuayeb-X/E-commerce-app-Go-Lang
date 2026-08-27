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

### 🏗️ Modular Architecture

The project separates different responsibilities into dedicated packages.

The architecture includes:

- Routes
- Handlers
- Middleware
- Database layer
- Utility functions
- Configuration

---

## 🛠️ Technology Stack

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

##  Create Product

Creates a new product.

**Method**

```http
POST /products
```

### Example Request

```json
{
  "name": "iPhone 15",
  "description": "Apple smartphone",
  "price": 120000,
  "img_url": "https://example.com/iphone.jpg"
}
```

### Example Response

```json
{
  "id": 1,
  "name": "iPhone 15",
  "description": "Apple smartphone",
  "price": 120000,
  "img_url": "https://example.com/iphone.jpg"
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
## 🗑️ Delete Product

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

---


