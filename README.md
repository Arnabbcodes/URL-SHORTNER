# 🔗 URL Shortener in Go

A simple **URL Shortener REST API built with Go (Golang)** that converts long URLs into short, unique identifiers and redirects users back to the original URL.

This project was created as a beginner-friendly backend project to understand how **Go HTTP servers, REST APIs, JSON, URL routing, hashing, and API testing** work together.

## 🚀 Features

- 🔗 Convert a long URL into a short URL ID
- ⚡ Generate short IDs using **MD5 hashing**
- 💾 Store shortened URLs in an in-memory Go map
- 🔄 Redirect short URLs to their original URLs
- 📡 REST API endpoints for URL shortening and redirection
- 📦 JSON request and response handling
- 🧪 API testing using **Postman**
- 🌐 Local HTTP server running on port `5000`
- ❌ Handles invalid requests and URLs that are not found

## 🛠️ Technologies Used

- **Go (Golang)** — Backend and HTTP server
- **net/http** — HTTP server and routing
- **encoding/json** — JSON request/response handling
- **crypto/md5** — Short ID generation
- **encoding/hex** — Converting the hash into a readable string
- **Postman** — API testing
- **In-memory Map** — Temporary URL storage

## 🔄 How It Works

The application follows this basic flow:

```text
Long URL
   ↓
MD5 Hash
   ↓
Hexadecimal Hash
   ↓
First 8 Characters
   ↓
Short URL ID
   ↓
Stored in Go Map
   ↓
Redirect using Short URL
```

For example:

```text
Original URL:
https://jsonplaceholder.typicode.com/todos/1

        ↓

Short ID:
a1b2c3d4

        ↓

Redirect URL:
http://localhost:5000/redirect/a1b2c3d4
```

## 📡 API Endpoints

### 1. Root Endpoint

**GET**

```text
http://localhost:5000/
```

Returns:

```text
Hello World!
```

### 2. Create Short URL

**POST**

```text
http://localhost:5000/shorten
```

Request body:

```json
{
    "url": "https://jsonplaceholder.typicode.com/todos/1"
}
```

Example response:

```json
{
    "short_url": "a1b2c3d4"
}
```

### 3. Redirect to Original URL

**GET**

```text
http://localhost:5000/redirect/a1b2c3d4
```

The server looks up the short ID and redirects the user to the original URL.

## 🧪 Testing with Postman

The API was tested using **Postman**.

The `/shorten` endpoint can be tested by sending a `POST` request with a JSON body containing the original URL.

Example:

```text
POST http://localhost:5000/shorten
```

Body → `raw` → `JSON`

```json
{
    "url": "https://example.com"
}
```

The server returns the generated short URL ID.

The returned ID can then be used with:

```text
GET http://localhost:5000/redirect/<short_id>
```

to test the redirection functionality.

## 📁 Project Structure

```text
URL-Shortener/
│
├── main.go
└── README.md
```

## ▶️ Running the Project

### 1. Clone the repository

```bash
git clone <your-repository-url>
```

### 2. Open the project directory

```bash
cd URL-Shortener
```

### 3. Run the Go server

```bash
go run main.go
```

The server will start at:

```text
http://localhost:5000
```

### 4. Test the API

Use **Postman** to send requests to:

```text
POST http://localhost:5000/shorten
```

and then use the generated ID with:

```text
GET http://localhost:5000/redirect/<short_id>
```

## ⚠️ Current Limitations

Shortened URLs are currently stored in an **in-memory map**.

Therefore:

- Data is lost when the server stops.
- There is no permanent database.
- Authentication is not implemented.
- MD5 is used for demonstration rather than security.
- Short IDs are limited to the first 8 characters of the generated hash.

## 🎯 Learning Objectives

- Building an HTTP server using Go
- Creating API endpoints
- Handling GET and POST requests
- Working with JSON in Go
- Using maps for temporary data storage
- Hash generation using MD5
- URL redirection
- HTTP status codes
- Error handling
- Testing APIs using Postman
- Understanding the basic structure of a backend REST API

## 🔮 Future Improvements

Possible improvements include:

- 🗄️ Add MySQL/PostgreSQL database storage
- 🔐 Add user authentication
- 📊 Add URL analytics and click tracking
- ⏳ Add URL expiration
- 🔑 Generate random short IDs
- 🌐 Add a frontend interface
- 🚀 Deploy the API online
- 📈 Add visit statistics for shortened URLs

---.
