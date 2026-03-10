# Go Echo Server

A minimal HTTP server built with Go that demonstrates how to handle POST requests, decode JSON bodies, and send JSON responses.

## Description

This project serves as a basic introduction to the `net/http` and `encoding/json` packages in Go. It listens for incoming messages and "echoes" them back to the client with a success status.

## API Specification

### Echo Message
Sends a text message to the server and receives the same message back.

**Endpoint:** `POST /echo`

**Request Body:**
```json
{
  "message": "Your text here"
}