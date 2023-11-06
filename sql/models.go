package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/thrillee/rssagg/internal/database"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Name     string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		Created:  dbUser.Created,
		Modified: dbUser.Modified,
		Name:     dbUser.Name,
	}
}
