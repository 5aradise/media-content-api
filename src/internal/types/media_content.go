package types

import (
	"time"
)

type MediaContent struct {
	Id          int32            `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description,omitempty"`
	Body        string           `json:"body"`
	ContentType MediaContentType `json:"content_type"`
	CreatedAt   time.Time        `json:"created_at"`
	UserID      int32            `json:"user_id"`
}
