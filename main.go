package main

import (
	"project-1/db"
	"project-1/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
