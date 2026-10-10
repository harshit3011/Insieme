package models

import "time"

type Like struct {
	ID           string
	LikedBy      string
	PostIdentity string
	LikedAt      time.Time
}
