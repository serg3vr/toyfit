package main

import (
	"toyfit/config"
	"toyfit/src/router"

	_ "github.com/lib/pq"
)

func main() {
	// config.DB = config.InitDB()
	config.InitDB()
	router.Start()
}
