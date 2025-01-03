package main

import (
	"car-rental-ums/cmd"
	"car-rental-ums/helpers"
)

func main() {
	// setup logger
	helpers.SetupLogger()

	// setup config
	helpers.SetupConfig()

	// setup db
	helpers.SetupDB()

	// setup server
	cmd.ServeHTTP()
}
