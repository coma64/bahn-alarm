package models

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"time"
)

type User struct {
	db.IdModel
	Name                         string
	PasswordHash                 string
	IsAdmin                      bool
	CreatedAt                    time.Time
	NotificationThresholdMinutes int
}

func (u *User) ToSchema() *server.User {
	return &server.User{
		CreatedAt: u.CreatedAt,
		Id:        u.Id,
		IsAdmin:   u.IsAdmin,
		Name:      u.Name,
	}
}
