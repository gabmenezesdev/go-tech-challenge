# go-tech-challenge

## Key Features

### 1. **Farm Creation**
- **POST** `/api/v1/farm`: Creates a new farm with details such as name, land area, unit, address, and crops.

### 2. **Farm Deletion**
- **DELETE** `/api/v1/farm/{id}`: Deletes a farm using its ID.

### 3. **Farm Retrieval**
- **GET** `/api/v1/farm`: Returns a list of farms with filters (name, unit, crop type, etc.) and pagination.

## Technical Stack

- **Backend**: Go (Golang) 1.23.4
- **Framework**: Gin-Gonic
- **Database**: MongoDB

## Project Structure

- **Controller Layer**: Defines routes and manages HTTP requests.
- **Use Case Layer**: Contains business logic for farm creation, deletion, and retrieval.
- **Repository Layer**: Interacts with the database (MongoDB) for data persistence and retrieval.
- **Domain Layer**: Defines domain models and related to farms.
- **DAO Layer**: Interacts with the database (MongoDB) for data persistence and retrieval.
- **Infra Layer**: Provides infrastructure-related components, such as database connections and http drivers.
- **Shared Layer**: Reusable functions, such as error responses and logging.

## Running Locally

### Option 1: Using Docker Compose

1. **Clone the repository**:
    ```bash
    git clone https://github.com/gabmenezesdev/go-tech-challenge.git
    cd go-tech-challenge
    ```

2. **Set up your environment variables**:
    ```bash
    cp .env.example .env
    ```
    - Adjust the values in `.env` as needed (e.g., MongoDB connection details).

3. **Build and start the application with Docker Compose**:
    ```bash
    docker-compose up --build
    ```
    or
    ```bash
    docker compose up --build
    ```

4. **Access the application**: Once the Docker containers are up and running, the application will be accessible on `http://localhost:3000`.

5. **Access the docs**: Access `http://localhost:3000/swagger/index.html` to see the docs and test the api.

6. **Stop the application**:
    ```bash
    docker-compose down
    ```
    or

    ```bash
    docker compose down
    ```

### Option 2: Without Docker Compose (Using .env.example)

1. **Clone the repository**:
    ```bash
    git clone https://github.com/gabmenezesdev/go-tech-challenge.git
    cd go-tech-challenge
    ```

2. **Set up your environment variables**:
    Copy `.env.example` to `.env`:
      ```bash
      cp .env.example .env
      ```
    - Adjust the values in `.env` as needed (e.g., MongoDB connection details).

3. **Start MongoDB**:
    Ensure that MongoDB is running locally or using a cloud service. If using Docker, you can run:
      ```bash
      docker run -d -p 27017:27017 --name mongodb mongo
      ```

4. **Run the application**:
    ```bash
    go run ./cmd/main.go
    ```

5. **Access the application**: The application will be available at `http://localhost:3000`.

6. **Access the docs**: Access `http://localhost:3000/swagger/index.html` to see the docs and test the api.

### Notes

- The application depends on MongoDB. If you haven't set up MongoDB locally, you can run it using Docker or connect to a remote MongoDB instance by adjusting the `.env` file.
- The default port for the application is `3000`, but it can be customized through environment variables.


## Deploy instructions

1. **Build the application**:
    - **Windows**:
    ```bash
    GOOS=windows GOARCH=amd64 go build -o ./build/app-windows.exe ./cmd/main.go
    ```
    - **Linux**:
    ```bash
    GOOS=linux GOARCH=amd64 go build -o ./build/app-linux ./cmd/main.go
    ```
    - **Mac**:
    ```bash
    GOOS=darwin GOARCH=amd64 go build -o ./build/app-mac ./cmd/main.go
    ```

2. **Place the `.env` file**:  
   After building the application, ensure that the `.env` file is in the same directory as the executable or in a directory where your application can access it. The `.env` file should contain any required environment variables for the application to run correctly.

    ```bash
    cp .env ./build/.env
    ```

3. **Execute the app**:  
   Run the built application from the command line. Make sure you are in the same directory as the executable or provide the correct path to the executable.

    - **Windows**:
    ```bash
    ./app-windows.exe
    ```
    - **Linux**:
    ```bash
    ./app-linux
    ```
    - **Mac**:
    ```bash
    ./app-mac
    ```

4. **Access the application**:  
   The application will be available at `http://localhost:3000`.


## How to run tests

## Prerequisites

- The application must be running.
- The database should be set up for integration tests.
- The -v flag provides detailed output for each test, including test results and any error messages.

## Running Tests

### Run All Tests

To run all tests (including unit and integration tests), use the following command:

```bash
docker-compose up --build
```

```bash
go test ./... -v
```

### Run Integration Tests

```bash
go test ./internal/application/use-case -v
```

### Run Unit Tests

```bash
go test ./internal/domain/... -v
```
