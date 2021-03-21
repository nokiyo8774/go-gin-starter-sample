package main

//go:generate sqlboiler --wipe mysql
import (
	"app/common/config"
	"app/common/db"
	"app/server"
)

func main() {
	config.Load()
	db.Open()
	defer db.Close()
	server.Init()
}
