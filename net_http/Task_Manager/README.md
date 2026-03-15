# Task Manager API 🚀

This is my second Go project. It is a RESTful API for managing a task list, designed to be **thread-safe** and handle concurrent requests correctly.

## 🧠 What I Learned
* **Advanced JSON**: Working with `json.NewDecoder` and `json.NewEncoder` to handle data streams.
* **Concurrency**: Implementing `sync.Mutex` to prevent **Data Races** during simultaneous read/write operations.
* **Resource Safety**: Using the `defer` keyword to ensure locks are always released (`Unlock`), preventing server deadlocks.
* **HTTP Standards**: Correctly separating logic between `GET` and `POST` methods and using appropriate HTTP status codes.

## 🛠 Tech Stack
* **Language**: Go (Golang)
* **Standard Library**: `net/http`, `encoding/json`, `sync`

## 📡 Endpoints

### 1. Get All Tasks
* **URL**: `/tasks`
* **Method**: `GET`
* **Description**: Returns the full list of tasks in JSON format.
* **Safety**: Uses a Mutex lock to ensure the list isn't being modified while it's being read.

### 2. Add a New Task
* **URL**: `/add`
* **Method**: `POST`
* **Headers**: `Content-Type: application/json`
* **Request Body**:
    ```json
    {
      "ID": 1,
      "Title": "Buy milk",
      "Done": false
    }
    ```
* **Safety**: Uses `mu.Lock()` before performing an `append` to protect the integrity of the slice.

## 🚦 How to Run
1. Clone the repository.
2. Run the server:
   ```bash
   go run main.go