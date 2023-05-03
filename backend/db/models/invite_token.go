package models

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"time"
)

type InviteToken struct {
	db.IdModel
	Token       string
	CreatedById int
	UsedById    *int
	ExpiresAt   time.Time
}
