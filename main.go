package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/sarangi/simplebank/api"
	db "github.com/sarangi/simplebank/db/sqlc"
	"github.com/sarangi/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
