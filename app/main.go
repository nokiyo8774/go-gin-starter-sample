package main

//go:generate sqlboiler --wipe mysql
import (
	config "app/common/config/db"
	"app/common/db"
	"app/routers"
)

func init() {
	config.Setup()
	db.Setup()
}

func main() {
	r := routers.Init()
	r.Run()
}
