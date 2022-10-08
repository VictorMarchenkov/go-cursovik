package main

import (
	"go-cursovic/database"
	"go-cursovic/routes"
)

func main() {
	database.ConnectToDB()
	routes.SetupAndListen()
}
