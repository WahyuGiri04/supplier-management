# Supplier Management API

A RESTful API service for managing suppliers, their groups, addresses, and contacts built with Go, Gin framework, and PostgreSQL.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Database Setup](#database-setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Request & Response Examples](#request--response-examples)
- [Project Structure](#project-structure)

## Features

- ✅ CRUD operations for Suppliers
- ✅ CRUD operations for Supplier Groups
- ✅ CRUD operations for Addresses
- ✅ CRUD operations for Contacts
- ✅ Pagination support
- ✅ Soft delete functionality
- ✅ Search by field and name
- ✅ Clean architecture with separation of concerns
- ✅ Generic base handlers, services, and repositories

## Tech Stack

- **Language**: Go 1.24.1
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM
- **Environment**: godotenv
- **UUID**: Google UUID

## Prerequisites

- Go 1.24.1 or higher
- PostgreSQL 12 or higher
- Git

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd supplier-management
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

## Configuration

1. **Create environment file**
   ```bash
   cp .env.example .env
   ```

2. **Configure environment variables in `.env`**
   ```env
   SERVICE_PORT=5001
   SERVER_PATH=/api/v1/supplier-service

   # Database configuration
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=dudung
   DB_PASSWORD=dudung123
   DB_NAME=dudungdb
   ```

## Database Setup

1. **Create PostgreSQL database**
   ```sql
   CREATE DATABASE dudungdb;
   ```

2. **Run database migration**
   ```bash
   psql -U dudung -d dudungdb -f db/db.sql
   ```

   This will create:
   - `supplier_management` schema
   - Required tables with proper relationships
   - Enable `pgcrypto` extension for UUID generation

## Running the Application

1. **Start the server**
   ```bash
   go run main.go
   ```

2. **Verify the server is running**
   ```bash
   curl http://localhost:5001/api/v1/supplier-service/suppliers
   ```

The API will be available at: `http://localhost:5001/api/v1/supplier-service`

## API Endpoints

### Base URL
```
http://localhost:5001/api/v1/supplier-service
```

### Suppliers

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/suppliers` | Get all suppliers |
| GET | `/suppliers/pagination?page=1&page_size=10` | Get suppliers with pagination |
| GET | `/suppliers/{uuid}` | Get supplier by UUID |
| POST | `/suppliers` | Create new supplier |
| PUT | `/suppliers/{uuid}` | Update supplier |

### Supplier Groups

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/supplier-groups` | Get all supplier groups |
| GET | `/supplier-groups/pagination?page=1&page_size=10` | Get supplier groups with pagination |
| GET | `/supplier-groups/{uuid}` | Get supplier group by UUID |
| POST | `/supplier-groups` | Create new supplier group |
| PUT | `/supplier-groups/{uuid}` | Update supplier group |

### Addresses

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/address` | Get all addresses |
| GET | `/address/pagination?page=1&page_size=10` | Get addresses with pagination |
| GET | `/address/{uuid}` | Get address by UUID |
| POST | `/address` | Create new address |
| PUT | `/address/{uuid}` | Update address |

### Contacts

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/contacts` | Get all contacts |
| GET | `/contacts/pagination?page=1&page_size=10` | Get contacts with pagination |
| GET | `/contacts/{uuid}` | Get contact by UUID |
| POST | `/contacts` | Create new contact |
| PUT | `/contacts/{uuid}` | Update contact |

## Request & Response Examples

### Create Supplier

**Request:**
```http
POST /api/v1/supplier-service/suppliers
Content-Type: application/json

{
  "code": "SUP001",
  "name": "ABC Company",
  "nickname": "ABC Corp",
  "logo": "https://example.com/logo.png"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Entity created successfully",
  "data": {
    "uuid": "123e4567-e89b-12d3-a456-426614174000",
    "code": "SUP001",
    "name": "ABC Company",
    "nickname": "ABC Corp",
    "logo": "https://example.com/logo.png",
    "is_active": true,
    "is_deleted": false,
    "supplierGroup": [],
    "address": [],
    "contact": []
  }
}
```

### Get All Suppliers

**Request:**
```http
GET /api/v1/supplier-service/suppliers
```

**Response:**
```json
{
  "status": 200,
  "message": "Entities retrieved successfully",
  "data": [
    {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "code": "SUP001",
      "name": "ABC Company",
      "nickname": "ABC Corp",
      "logo": "https://example.com/logo.png",
      "is_active": true,
      "is_deleted": false,
      "supplierGroup": [
        {
          "uuid": "456e7890-e89b-12d3-a456-426614174001",
          "supplierId": 1,
          "groupName": "Electronics",
          "is_active": true,
          "is_deleted": false
        }
      ],
      "address": [
        {
          "uuid": "789e0123-e89b-12d3-a456-426614174002",
          "supplierId": 1,
          "name": "Head Office",
          "fullAddress": "123 Main St, City, Country",
          "isMain": true,
          "is_active": true,
          "is_deleted": false
        }
      ],
      "contact": [
        {
          "uuid": "012e3456-e89b-12d3-a456-426614174003",
          "supplierId": 1,
          "name": "John Doe",
          "jobPosition": "Manager",
          "email": "john.doe@abc.com",
          "phoneNumber": "+1234567890",
          "mobilePhoneNumber": "+0987654321",
          "isMain": true,
          "is_active": true,
          "is_deleted": false
        }
      ]
    }
  ]
}
```

### Get Suppliers with Pagination

**Request:**
```http
GET /api/v1/supplier-service/suppliers/pagination?page=1&page_size=5
```

**Response:**
```json
{
  "status": 200,
  "message": "Paginated entities retrieved successfully",
  "data": {
    "page": 1,
    "page_size": 5,
    "total_rows": 25,
    "total_pages": 5,
    "data": [
      {
        "uuid": "123e4567-e89b-12d3-a456-426614174000",
        "code": "SUP001",
        "name": "ABC Company",
        "nickname": "ABC Corp",
        "logo": "https://example.com/logo.png",
        "is_active": true,
        "is_deleted": false
      }
    ]
  }
}
```

### Create Supplier Group

**Request:**
```http
POST /api/v1/supplier-service/supplier-groups
Content-Type: application/json

{
  "supplierId": 1,
  "groupName": "Electronics"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Entity created successfully",
  "data": {
    "uuid": "456e7890-e89b-12d3-a456-426614174001",
    "supplierId": 1,
    "groupName": "Electronics",
    "is_active": true,
    "is_deleted": false
  }
}
```

### Create Address

**Request:**
```http
POST /api/v1/supplier-service/address
Content-Type: application/json

{
  "supplierId": 1,
  "name": "Head Office",
  "fullAddress": "123 Main St, City, Country",
  "isMain": true
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Entity created successfully",
  "data": {
    "uuid": "789e0123-e89b-12d3-a456-426614174002",
    "supplierId": 1,
    "name": "Head Office",
    "fullAddress": "123 Main St, City, Country",
    "isMain": true,
    "is_active": true,
    "is_deleted": false
  }
}
```

### Create Contact

**Request:**
```http
POST /api/v1/supplier-service/contacts
Content-Type: application/json

{
  "supplierId": 1,
  "name": "John Doe",
  "jobPosition": "Manager",
  "email": "john.doe@abc.com",
  "phoneNumber": "+1234567890",
  "mobilePhoneNumber": "+0987654321",
  "isMain": true
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Entity created successfully",
  "data": {
    "uuid": "012e3456-e89b-12d3-a456-426614174003",
    "supplierId": 1,
    "name": "John Doe",
    "jobPosition": "Manager",
    "email": "john.doe@abc.com",
    "phoneNumber": "+1234567890",
    "mobilePhoneNumber": "+0987654321",
    "isMain": true,
    "is_active": true,
    "is_deleted": false
  }
}
```

### Error Response Example

**Response:**
```json
{
  "status": 400,
  "message": "Invalid request body: Key: 'Supplier.Name' Error:Field validation for 'Name' failed on the 'required' tag",
  "data": null
}
```

## Project Structure

```
supplier-management/
├── config/
│   └── database.go          # Database configuration
├── db/
│   └── db.sql              # Database schema and migration
├── handler/
│   ├── base/
│   │   └── base_handler.go # Generic CRUD handlers
│   ├── address_handler.go   # Address-specific handlers
│   ├── contacts_handler.go  # Contacts-specific handlers
│   ├── supplier_group_handler.go # Supplier group handlers
│   └── supplier_handler.go  # Supplier-specific handlers
├── model/
│   ├── base/
│   │   ├── base_model.go   # Base model with common fields
│   │   ├── base_response.go # Standard API response structure
│   │   └── pagination.go   # Pagination structure
│   ├── address.go          # Address model
│   ├── contacts.go         # Contact model
│   ├── supplier.go         # Supplier model with relationships
│   └── supplier_group.go   # Supplier group model
├── repo/
│   ├── base/
│   │   └── base_repository.go # Generic CRUD repository
│   ├── address_repo.go     # Address repository
│   ├── contacts_repo.go    # Contacts repository
│   ├── supplier_group_repo.go # Supplier group repository
│   └── supplier_repo.go    # Supplier repository
├── router/
│   └── routes.go           # API route definitions
├── service/
│   ├── base/
│   │   └── base_service.go # Generic business logic
│   ├── address_service.go  # Address business logic
│   ├── contacts_service.go # Contacts business logic
│   ├── supplier_group_service.go # Supplier group logic
│   └── supplier_service.go # Supplier business logic
├── util/
│   └── response_helper.go  # HTTP response helpers
├── .env                    # Environment variables
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── main.go                 # Application entry point
└── README.md              # This file
```

## Development

### Adding New Entities

1. Create model in `model/` directory
2. Create repository in `repo/` directory (extend base repository)
3. Create service in `service/` directory (extend base service)
4. Create handler in `handler/` directory (extend base handler)
5. Add routes in `router/routes.go`

### Database Migration

When adding new fields or tables, update the `db/db.sql` file and run:

```bash
psql -U dudung -d dudungdb -f db/db.sql
```