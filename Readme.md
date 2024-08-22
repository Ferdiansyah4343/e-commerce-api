# E-Commerce API Documentation

## Table of Contents

1. [Overview](#overview)
2. [Base URL](#base-url)
3. [Endpoints](#endpoints)
    - [1. Get All Products](#1-get-all-products)
    - [2. Get Product by ID](#2-get-product-by-id)
    - [3. Create Product](#3-create-product)
    - [4. Update Product](#4-update-product)
    - [5. Delete Product](#5-delete-product)
    - [6. Get All Users](#6-get-all-users)
    - [7. Get User by ID](#7-get-user-by-id)
    - [8. Create User](#8-create-user)
    - [9. Update User](#9-update-user)
    - [10. Delete User](#10-delete-user)
    - [11. Create Order](#11-create-order)
    - [12. Get Order by ID](#12-get-order-by-id)
    - [13. Get All Orders](#13-get-all-orders)
4. [Error Handling](#error-handling)
5. [Response Format](#response-format)

## Overview

This API provides endpoints for managing products, users, and orders in an e-commerce application. It uses **Gin** as the web framework, **GORM** as the ORM for database interactions, and **MySQL** as the database system.

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Get All Products

**Endpoint:** `GET /products`

**Description:** Retrieve a list of all products.

**Response:**

- **Success (200 OK):**

    ```json
    [
        {
            "id": 1,
            "name": "Product Name",
            "description": "Product Description",
            "price": 99.99,
            "stock": 10,
            "created_at": "2024-08-23T10:00:00Z",
            "updated_at": "2024-08-23T10:00:00Z"
        }
    ]
    ```

- **Error (500 Internal Server Error):**

    ```json
    {
        "error": "Failed to retrieve products."
    }
    ```

### 2. Get Product by ID

**Endpoint:** `GET /products/:id`

**Description:** Retrieve a single product by its ID.

**Parameters:**

- `id` (path): Product ID (integer)

**Response:**

- **Success (200 OK):**

    ```json
    {
        "id": 1,
        "name": "Product Name",
        "description": "Product Description",
        "price": 99.99,
        "stock": 10,
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "Product not found."
    }
    ```

### 3. Create Product

**Endpoint:** `POST /products`

**Description:** Create a new product.

**Request Body:**

```json
{
    "name": "New Product",
    "description": "Product Description",
    "price": 49.99,
    "stock": 20
}
```

**Response:**

- **Success (201 Created):**

    ```json
    {
        "id": 2,
        "name": "New Product",
        "description": "Product Description",
        "price": 49.99,
        "stock": 20,
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (400 Bad Request):**

    ```json
    {
        "error": "Invalid input data."
    }
    ```

### 4. Update Product

**Endpoint:** `PUT /products/:id`

**Description:** Update an existing product by its ID.

**Parameters:**

- `id` (path): Product ID (integer)

**Request Body:**

```json
{
    "name": "Updated Product",
    "description": "Updated Description",
    "price": 59.99,
    "stock": 15
}
```

**Response:**

- **Success (200 OK):**

    ```json
    {
        "id": 1,
        "name": "Updated Product",
        "description": "Updated Description",
        "price": 59.99,
        "stock": 15,
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "Product not found."
    }
    ```

### 5. Delete Product

**Endpoint:** `DELETE /products/:id`

**Description:** Delete a product by its ID.

**Parameters:**

- `id` (path): Product ID (integer)

**Response:**

- **Success (204 No Content):**

    ```json
    {}
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "Product not found."
    }
    ```

### 6. Get All Users

**Endpoint:** `GET /users`

**Description:** Retrieve a list of all users.

**Response:**

- **Success (200 OK):**

    ```json
    [
        {
            "id": 1,
            "name": "User Name",
            "email": "user@example.com",
            "address": "User Address",
            "phone": "1234567890",
            "created_at": "2024-08-23T10:00:00Z",
            "updated_at": "2024-08-23T10:00:00Z"
        }
    ]
    ```

- **Error (500 Internal Server Error):**

    ```json
    {
        "error": "Failed to retrieve users."
    }
    ```

### 7. Get User by ID

**Endpoint:** `GET /users/:id`

**Description:** Retrieve a single user by their ID.

**Parameters:**

- `id` (path): User ID (integer)

**Response:**

- **Success (200 OK):**

    ```json
    {
        "id": 1,
        "name": "User Name",
        "email": "user@example.com",
        "address": "User Address",
        "phone": "1234567890",
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "User not found."
    }
    ```

### 8. Create User

**Endpoint:** `POST /users`

**Description:** Create a new user.

**Request Body:**

```json
{
    "name": "New User",
    "email": "newuser@example.com",
    "password": "password",
    "address": "User Address",
    "phone": "0987654321"
}
```

**Response:**

- **Success (201 Created):**

    ```json
    {
        "id": 2,
        "name": "New User",
        "email": "newuser@example.com",
        "address": "User Address",
        "phone": "0987654321",
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (400 Bad Request):**

    ```json
    {
        "error": "Invalid input data."
    }
    ```

### 9. Update User

**Endpoint:** `PUT /users/:id`

**Description:** Update an existing user by their ID.

**Parameters:**

- `id` (path): User ID (integer)

**Request Body:**

```json
{
    "name": "Updated User",
    "email": "updateduser@example.com",
    "address": "Updated Address",
    "phone": "1122334455"
}
```

**Response:**

- **Success (200 OK):**

    ```json
    {
        "id": 1,
        "name": "Updated User",
        "email": "updateduser@example.com",
        "address": "Updated Address",
        "phone": "1122334455",
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z"
    }
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "User not found."
    }
    ```

### 10. Delete User

**Endpoint:** `DELETE /users/:id`

**Description:** Delete a user by their ID.

**Parameters:**

- `id` (path): User ID (integer)

**Response:**

- **Success (204 No Content):**

    ```json


    {}
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "User not found."
    }
    ```

### 11. Create Order

**Endpoint:** `POST /orders`

**Description:** Create a new order.

**Request Body:**

```json
{
    "user_id": 1,
    "total": 199.98,
    "status": "Pending",
    "items": [
        {
            "product_id": 1,
            "quantity": 2,
            "price": 99.99
        }
    ]
}
```

**Response:**

- **Success (201 Created):**

    ```json
    {
        "id": 1,
        "user_id": 1,
        "total": 199.98,
        "status": "Pending",
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z",
        "items": [
            {
                "id": 1,
                "order_id": 1,
                "product_id": 1,
                "quantity": 2,
                "price": 99.99,
                "total": 199.98
            }
        ]
    }
    ```

- **Error (400 Bad Request):**

    ```json
    {
        "error": "Invalid input data."
    }
    ```

### 12. Get Order by ID

**Endpoint:** `GET /orders/:id`

**Description:** Retrieve a single order by its ID.

**Parameters:**

- `id` (path): Order ID (integer)

**Response:**

- **Success (200 OK):**

    ```json
    {
        "id": 1,
        "user_id": 1,
        "total": 199.98,
        "status": "Pending",
        "created_at": "2024-08-23T10:00:00Z",
        "updated_at": "2024-08-23T10:00:00Z",
        "items": [
            {
                "id": 1,
                "order_id": 1,
                "product_id": 1,
                "quantity": 2,
                "price": 99.99,
                "total": 199.98
            }
        ]
    }
    ```

- **Error (404 Not Found):**

    ```json
    {
        "error": "Order not found."
    }
    ```

### 13. Get All Orders

**Endpoint:** `GET /orders`

**Description:** Retrieve a list of all orders.

**Response:**

- **Success (200 OK):**

    ```json
    [
        {
            "id": 1,
            "user_id": 1,
            "total": 199.98,
            "status": "Pending",
            "created_at": "2024-08-23T10:00:00Z",
            "updated_at": "2024-08-23T10:00:00Z",
            "items": [
                {
                    "id": 1,
                    "order_id": 1,
                    "product_id": 1,
                    "quantity": 2,
                    "price": 99.99,
                    "total": 199.98
                }
            ]
        }
    ]
    ```

- **Error (500 Internal Server Error):**

    ```json
    {
        "error": "Failed to retrieve orders."
    }
    ```

## Error Handling

All error responses will include an `error` field with a description of the issue. The following HTTP status codes are used for errors:

- **400 Bad Request:** Invalid input data.
- **404 Not Found:** Resource not found.
- **500 Internal Server Error:** Unexpected error occurred on the server.

## Response Format

All responses are in JSON format and include the appropriate HTTP status code. Successful responses will contain the requested data, while error responses will provide a description of the problem.
