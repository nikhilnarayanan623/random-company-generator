# Random Company Generator Microservices

Random Company Generator is a Go microservices project designed to generate random company details based on user input. It consists of an API Gateway, an Authentication Service, and a Company Service, each serving a specific purpose in the system.

## Overview

The Random Company Generator project is built around the concept of generating random company details such as company name, CEO name, and the total number of employees based on user input. It is a microservices architecture with the following components:

- **API Gateway:** Utilizes the Gin web framework to route and handle incoming requests. It communicates with the Authentication Service and the Company Service to manage user authentication and retrieve generated company details. Data is streamed from the Company Service to the API Gateway using gRPC.

- **Authentication Service:** Manages user authentication, including sign-up, sign-in, and JWT token verification. It uses a PostgreSQL database to store user information securely.

- **Company Service:** Generates random company details based on user input and stores this information in a MongoDB database. It communicates with the API Gateway via gRPC and streams generated company data to the API Gateway for retrieval.

## Requirements

Before getting started, ensure you have the following requirements in place:

- Go programming environment
- PostgreSQL database for the Authentication Service
- MongoDB for the Company Service
- gRPC for communication between services
- Gin web framework for the API Gateway

## Getting Started

### Prerequisites

To use this boilerplate, you need to have the following installed on your system:

- Go programming environment
- PostgreSQL database for the Authentication Service
- MongoDB for the Company Service
- gRPC for communication between services
- Gin web framework for the API Gateway

### Installation

  #### 1. Clone the repository:

  ```bash
    git clone https://github.com/nikhilnarayanan623/random-company-generator.git && \
    cd ./random-company-generator
  ```

### Instructions

  #### 1. API Gateway
  ##### 1. Install dependencies
  ```bash
    ## Assuming you are in root of the project
    cd ./api-gateway && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
     ## api gateway
     API_PORT="port that you want to run api gateway service"
     ## auth service
     AUTH_SERVICE_HOST="auth service host"
     AUTH_SERVICE_PORT="auth service port"
     ## company service
     COMPANY_SERVICE_HOST="running company service host"
     COMPANY_SERVICE_PORT="running company service port"
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```
  #### 2. Auth Service
  ##### 1. Install dependencies
  ```bash
    ## Assuming you are in root of the project
    cd ./auth-service && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
    AUTH_SERVICE_HOST="auth service host"
    AUTH_SERVICE_PORT="auth service port"
    ## Database
    DB_HOST="database running host"
    DB_PORT="database running port"
    DB_NAME="database name"
    DB_USER="database user"
    DB_PASSWORD="database user password"
    ## JWT
    JWT_KEY="a key that you wan't to sign in for token"
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```
  #### 3. Company Service
  ##### 1. Install dependencies
  ```bash
    cd ./company-service && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
    COMPANY_SERVICE_HOST="company service host"
    COMPANY_SERVICE_PORT="company service port"
    ## database
    DB_PROTOCOL="mongo db protocol"
    DB_HOST="mongo host"
    DB_NAME="database name"
    DB_USER="databas user"
    DB_PASSWORD="database password"
    DB_OPTIONS="databse options"   put the options in ''
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```

## Usage

## Usage
### 1. Live API Documentation (Pending..)
If you are running the project then visit (http://localhost:{$API_PORT}/swagger/index.html)

### Authentication Service

1. Sign Up: Create a new user account by providing the necessary details.
2. Sign In: Authenticate with your credentials and receive a JWT token.
3. JWT Token Verification: Use the JWT token for secure access to protected routes and services.

### Company Service

1. Configure the Company Service envs,.
2. Start the Company Service.
3. Access gRPC methods to generate random company details based on user input.

### API Gateway

1. Handle incoming HTTP requests and route them to the appropriate service.
2. Stream generated company data from the Company Service to the client using gRPC.
3. Secure endpoints using JWT token verification.

## Contributing

We welcome contributions to this project. To contribute, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and submit a pull request.
4. Ensure your code follows the project's coding standards.

