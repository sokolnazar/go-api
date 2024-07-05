package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sokolnazar/go-api/cmd/api"
	"github.com/sokolnazar/go-api/config"
	"github.com/sokolnazar/go-api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.NewConfig{
		User:                 config.Envs.DBUser,
		Password:             config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBname:               config.Envs.DBName,
		Net:                  "ecom",
		AllowNativePasswoord: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
