package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/venu-prasath/go-bank/api"
	db "github.com/venu-prasath/go-bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secretPass123@localhost:5432/go-bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}