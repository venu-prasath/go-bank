package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/venu-prasath/go-bank/api"
	db "github.com/venu-prasath/go-bank/db/sqlc"
)


func main() {
	godotenv.Load("app.env")
	db_driver := os.Getenv("DB_DRIVER")
	db_source := os.Getenv("DB_SOURCE")
	serv_addr := os.Getenv("SERVER_ADDRESS")


	conn, err := sql.Open(db_driver, db_source)
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serv_addr)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}