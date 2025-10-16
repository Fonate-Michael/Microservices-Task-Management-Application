# API Documentation

## Authentication

All protected endpoints require JWT authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## API Endpoints

### User Service Endpoints

#### Register User

**Endpoint:** `POST /users/register`

**Description:** Register a new user account

**Request Body:**
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "message": "User registered successfully"
}
```

**Status Codes:**
- `200` - Success
- `400` - Bad request (invalid JSON)
- `500` - Internal server error

#### Login User

**Endpoint:** `POST /users/login`

**Description:** Authenticate user and receive JWT token

**Request Body:**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "token": "jwt-token-string"
}
```

**Status Codes:**
- `200` - Success
- `400` - Bad request (invalid JSON)
- `500` - Internal server error

#### Health Check

**Endpoint:** `GET /users/health`

**Description:** Check if user service is running

**Response:**
```json
{
  "message": "Bitch am alive!"
}
```

### Task Service Endpoints

#### Get All Tasks

**Endpoint:** `GET /task/tasks`

**Description:** Retrieve all tasks (public endpoint)

**Response:**
```json
{
  "tasks": [
    {
      "id": "integer",
      "user_id": "integer",
      "title": "string",
      "description": "string"
    }
  ]
}
```

**Status Codes:**
- `200` - Success
- `500` - Internal server error

#### Create Task

**Endpoint:** `POST /task/task`

**Description:** Create a new task (protected endpoint)

**Headers:**
```
Authorization: Bearer <jwt-token>
```

**Request Body:**
```json
{
  "title": "string",
  "description": "string"
}
```

**Response:**
```json
{
  "message": "Tasks Created successfully!"
}
```

**Status Codes:**
- `201` - Created
- `400` - Bad request (invalid JSON or missing token)
- `401` - Unauthorized (invalid or missing token)
- `500` - Internal server error

## Data Models

### User Model
```json
{
  "id": "integer",
  "username": "string",
  "email": "string",
  "password": "string (hashed)"
}
```

### Task Model
```json
{
  "id": "integer",
  "user_id": "integer (foreign key to users.id)",
  "title": "string",
  "description": "string"
}
```

## Error Handling

All endpoints follow consistent error response formats:

### Bad Request
```json
{
  "error": "Failed to bind json check ur json again"
}
```

### Unauthorized
```json
{
  "message": "Invalid Token"
}
```

### Internal Server Error
```json
{
  "error": "Failed to register user"
}
```

## Request/Response Examples

### User Registration Example

**Request:**
```bash
curl -X POST http://localhost:8002/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

**Response:**
```json
{
  "message": "User registered successfully"
}
```

### User Login Example

**Request:**
```bash
curl -X POST http://localhost:8002/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Create Task Example

**Request:**
```bash
curl -X POST http://localhost:8002/task/task \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete API documentation",
    "description": "Finish writing the API documentation"
  }'
```

**Response:**
```json
{
  "message": "Tasks Created successfully!"
}
```

## Database Schema

```sql
-- Users table
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Tasks table
CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(200) NOT NULL,
    description TEXT
);
```
