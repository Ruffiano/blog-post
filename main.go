package main

import (
	"database/sql"
	"log"
	"ruffiano/blog-post/api"

	_ "github.com/lib/pq"

	db "ruffiano/blog-post/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/blog_post?sslmode=disable"
	serverAddress = "0.0.0.0:4040"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create sever:", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
