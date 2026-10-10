package models

import "time"

type Comment struct {
	ID          string
	CommentedBy string
	PostId      string
	ParentId 	*string
	Image       string
	Content     string
	CommentedAt time.Time
}
