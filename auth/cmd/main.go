package main

import (
	"auth/api/server"
	"auth/internal/db"
)

func main() {
	dbClient := db.NewClient()

	users := db.NewUsers(dbClient)

	s := server.New(users)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
