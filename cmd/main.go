package main

import (
	"das/api/server"
	"das/internal/db"
)

func main() {
	dbClient := db.NewClient()

	polls := db.NewPolls(dbClient)

	s := server.New(polls)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
