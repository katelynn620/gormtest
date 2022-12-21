package main

import "gormtest/pkg/postgres"

func main() {
	db := postgres.DBClient{}

	db.Connect()
	defer db.Disconnect()
}
