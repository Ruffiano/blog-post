package main

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/ruffiano/blog-post/api"
	"github.com/ruffiano/blog-post/util"

	_ "github.com/lib/pq"

	db "github.com/ruffiano/blog-post/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config: ")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db:")
	}

	// runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	runGinServer(config, store)
}

// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal().Msg("cannot create new migrate instance: ")
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal().Msg("failed to run migrate up: ")
// 	}

// 	log.Print("db migrated successfully")
// }

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create sever: ")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server: ")
	}
}
