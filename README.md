## What Is Cats Social?
Cats Social is an application where cat owners can match their cats with each other.

## Key Technologies

1. **Go**: Go (or Golang) is the programming language that can develop an API with high performance and scalable.

2. **Gorilla Mux**: Gorilla Mux is a popular HTTP router for Go. It is used for routing incoming HTTP requests to the appropriate handler functions.

3. **PostgreSQL**: PostgreSQL is a powerful, open-source relational database management system. It is used for storing and managing data related to products, customers, orders, and more in the online store application.

4. **JWT (JSON Web Tokens)**: JWT is a standard for securely transmitting information between parties as a JSON object. In the online store application, JWT is used for implementing authentication and authorization mechanisms.

5. **Pgx**: Pgx is a pure Go driver and toolkit for PostgreSQL. The pgx driver is a low-level, high performance interface that exposes PostgreSQL-specific features such as LISTEN / NOTIFY and COPY. It also includes an adapter for the standard database/sql interface.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/agusheryanto182/go-cats-social.git
   ```

2. Go to folder go-cats-social

   ```bash
   cd go-cats-social
   ```

3. Set up env, copy the code and then paste it in your terminal

4. Create a new database

5. Copy this and paste it in terminal for migrate database

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
   ```

7. ```bash
   go run .
   ```

## Documentation

1. Postman

   ```bash
   https://documenter.getpostman.com/view/32137512/2sA3JFCQWq
   ```

2. Database

   ```bash
   https://dbdiagram.io/d/cats-social-662e1e645b24a634d0fd02c6
   ```
