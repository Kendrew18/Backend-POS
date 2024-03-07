package main

import (
	"Bakend-POS/db"
	"Bakend-POS/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":3333"))
}
