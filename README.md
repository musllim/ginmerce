## Ginmerce

Ginmerce is an e-commerce platform built with Go and Gin framework.

### Getting Started

Follow these instructions to set up and run the project on your local machine.

### Prerequisites

- Go (version 1.16 or higher)
- Git

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/musllim/ginmerce.git
    ```

2. Navigate to the project directory:
    ```sh
    cd ginmerce
    ```

3. Copy the example environment variables file and update it with your values:
    ```sh
    cp .env.example .env
    ```

### Running the Application

1. Initialize Swagger documentation:
    ```sh
    swag init
    swag init -g controllers/cart.go
    swag init -g controllers/users.go
    swag init -g controllers/products.go
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

3. Access the Swagger documentation at:
    ```
    {protocol}://{host}:{port}/swagger/index.html
    ```

### Contributing

Contributions are welcome! Please fork the repository and create a pull request.
