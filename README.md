### Go Banks

This is an experimental application to explore backend development with the following tech stack.

1. Golang - Programming language
2. Postgres - Database
3. Redis - In-memory cache
4. Gin - API framework
5. gRPC - Communication protocol
6. Docker - Containerization

#### Steps to build the project

1. Install the necessary build tools and libraries.
   - Golang
   - VSCode
   - Docker
   - sqlc
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

## Tasks

#### 1. Create migrations

Run the migration command mentioned above to initialize the migration files.

##### Installation

```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
# or
sudo apt-get install migrate

```

Check the migrate version

```
migrate -v
```

##### setup migration for the first time

1. Create a db/migration folder. This will hold all the versioned migration files.
2. setup migration by running:
   `migrate create -ext sql -dir db/migration -seq init_schema`
3. Add the schema information from dbdiagram.io to the init_schema_up file.
4. Run migration by executing:
   `migrate -path db/migration -database "<database-host-path> -verbose up`
5. For ease of use, add the previous run migration command to Makefile

#### 2. Setup sqlc

1. After installing sqlc, make sure to refer the sqlc's github page to configure sqlc.yaml script. You can refer the sqlc.yaml script in this project for a small database of 3 tables.
2. Make sure to update Makefile to include the `sqlc generate` command.
3. Create the necessary queries in the db/query directory and run `sqlc generate`
4. Check the generated db files for errors.

#### 3. Create Unit Tests for Database CRUD

Setup main_test.go and other test files for all the generated go files.
To test the setup of main_test.go, run the following command:

```
go test -timeout 30s github.com/venu-prasath/go-bank/db/sqlc -run ^TestMain$
```
