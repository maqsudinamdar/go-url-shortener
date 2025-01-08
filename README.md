# URL Shortener

This is a simple URL shortener service that allows users to shorten long URLs into short, easy-to-share links. The service uses CRC32 hashing for generating short URLs and stores them in a MySQL database.

## Features

- Generate short URLs from long URLs.
- Retrieve the original long URL using the short URL.
- Store short URL data in a MySQL database.

## Tech Stack

- **Go**: The backend is written in Go (Golang).
- **MySQL**: The short URLs and corresponding long URLs are stored in a MySQL database.
- **Gin** or **Chi**: The HTTP routing is handled by the Gin or Chi web framework.

## Setup Instructions

### Prerequisites

- Go (Golang) installed on your machine.
- MySQL server running and accessible.

### Clone the repository

```bash
git clone https://github.com/maqsudinamdar/go-url-shortener.git
cd go-url-shortener
```

## Install Dependencies
```bash
go mod tidy
```


### Set up the Database

```sql
CREATE DATABASE url_shortener;
USE url_shortener;

CREATE TABLE url (
    id INT AUTO_INCREMENT PRIMARY KEY,
    long_url VARCHAR(100),
    short_url VARCHAR(6)
);

```

### Run the Application

```bash
go run ./cmd/url-shortener
```
