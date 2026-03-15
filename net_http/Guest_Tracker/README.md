# Guest Tracker API 📝

This is a practice project in Go for learning the `net/http` package.
The server can register guests via JSON and dynamically greet them using the URL path.

## 🚀 Features
* **Guest Registration**: Accepts a name in JSON format at the `/enter` endpoint.
* **Dynamic Echo**: A catch-all handler that reads any path from the URL.
* **Security**: HTTP method validation and mandatory JSON header check.
* **Routing**: Uses a custom `http.ServeMux` for organized routing.

## 🛠 Tech Stack
* **Language**: Go (Golang)
* **Protocol**: HTTP
* **Logic**: Path handling via `r.URL.Path`.

## 📡 Endpoints

### 1. Register Guest (`/enter`)
* **Method**: `POST`
* **Body**: `{"name": "Minix"}`
* **Response**: Registers the guest in the system.

### 2. Dynamic Greeting (`/*`)
* **Method**: ANY
* **Description**: Uses `r.URL.Path` to greet the user.
* **Example**: Accessing `localhost:8080/Master` will return: "You are here: /Master".