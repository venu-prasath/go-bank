### Go Banks

This is an experimental application to explore backend development with the following tech stack.

1. Golang - Programming language
2. Postgres - Database
3. Redis - In-memory cache
4. Gin - API framework
5. gRPC - Communication protocol
6. Docker - Containerization

#### Steps to build the project

1. Install the necessary build tools.
   - Golang
   - VSCode
   - Docker
2. Install and run Postgres database inside Docker
   - Install Postgres
   - Install PgAdmin4
3. Create a simple database schema using dbdiagram.io
4. Execute the initial schema to load the tables
   - Install go migrate library as a standalone cli
   ```
     migrate create -ext sql -dir db/migration -seq init_schema
   ```
   The above command creates 2 files with different suffix (Up/Down Migration)

All the necessary scripts can be found in the Makefile for quick dev setup.
