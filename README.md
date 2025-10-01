# Coin Fetcher App

A Go-based service for fetching and storing cryptocurrency prices using the CoinGecko API, built with a **Hexagonal Architecture** (Ports and Adapters). The application supports multiple cryptocurrencies (e.g., Bitcoin, Ethereum) by their coin (e.g., `bitcoin`, `ethereum`), stores price data in a Redis then PostgreSQL database, and provides RESTful endpoints to retrieve prices and supported coins. It uses Docker for containerization and includes context-based timeout handling for API and database operations.

## Features
- Fetch real-time cryptocurrency prices from the CoinGecko API using coins (e.g., `bitcoin`, `ethereum`).
- You can change scheduler timer or coin name from .env
- Store price data in a PostgreSQL database with fields for coin ID, price, and timestamp.
- Retrieve the latest price or average price for a given coin over specified intervals (1min, 5min, 1day).
- Use Redis to save coin's price
- Containerized deployment using Docker.

## Project Structure
The project follows a Hexagonal Architecture to ensure modularity and testability:
```
coin-fetcher-app/
├── cmd/
│   └── server/
│       └── main.go           # Entry point of the application
├── internal/
│   ├── adapters/
│   │   ├── external/        # External service adapters (e.g., CoinGecko API)
│   │   ├── http/            # HTTP handlers and routes
│   │   ├── repository/      # Database repository (PostgreSQL with GORM AND REDIS)
│   │   └── scheduler/       # Background price fetching scheduler
│   ├── application/         # Business logic (PriceService)
│   └── domain/              # Domain models and interfaces
├── Dockerfile               # Docker configuration for building and running
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
└── README.md                # Project documentation
```

## Prerequisites
- **Go**: Version 1.22 or higher
- **Docker**: For containerized deployment
- **PostgreSQL**: For storing price data
- **REDIS**: For storing price data
- **Git**: For cloning the repository
- Access to the CoinGecko API (no authentication required for public endpoints)

## Installation

### Clone the Repository
```bash
git clone https://github.com/rzmsv/coin-fetcher-app.git
cd coin-fetcher-app
```


### Build and Run Locally
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Build the application:
   ```bash
   go run ./cmd/server
   ```

### Build and Run with Docker
*** USE PROXY ***

1. Build the Docker image:
   ```bash
   docker-compose build --no-cache
   ```
2. Run 
   ```bash
   docker-compose down (If you have already run it )
   ```
3. Run 
   ```bash
   docker-compose up
   ```

## Usage
The service exposes the following RESTful endpoints:
- **GET /api/price/history/:coin**: Retrieve the latest price (for `1min`) or average price (for `5min`, `1day`) for a given coin (e.g., `bitcoin`, `ethereum`).
  - Example: `curl http://localhost:8080/api/price/history/bitcoin`
  - Response: `{"price": 65000.12345678}`

  NOTICE: You want 2 APIs for get hestory and latest price (I know ) but I write all of this in one API for axample if you dont add interval you get latest price of coin from database (1min ago) and if you add interval=5min you get average price of 5min ******


## Database Schema
The `Coin` table is automatically created by GORM with the following schema:
```sql
	ID        uint
	Price     float64
	Coin      string
	Timestamp time.Time

## Contact
For questions or support, contact [rmussavi@gmail.com].