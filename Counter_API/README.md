# Go Counter API

A simple RESTful API built with Go standard library (`net/http`) to practice HTTP methods, JSON decoding, and state management in memory.

## Features

- **Get current count**: Retrieve the value of a global counter.
- **Add to count**: Send a JSON payload to increment the counter.
- **Reset count**: Reset the counter back to zero.

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/get` | Returns the current value of the counter. |
| `POST` | `/add` | Increments the counter by a specified amount. |
| `POST` | `/reset` | Resets the counter to 0. |

## Usage Examples

### 1. Get current value
**Request:**
```bash
GET http://localhost:9091/get