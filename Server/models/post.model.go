package models

import "time"

type Post struct {
	ID        string
	PostedBy  string
	Image     *string
	Caption   string
	Shares    int64
	Bookmarks int64
	CreatedAt time.Time
}
