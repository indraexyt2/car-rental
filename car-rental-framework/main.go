package main

import (
	"car-rental-framework/cmd"
	"car-rental-framework/helpers"
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
