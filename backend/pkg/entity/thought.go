package entity

import (
	"time"
)

type ThoughtInput struct {
	Passphrase string    `json:"passphrase" db:"passphrase" validate:"required"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtResponse struct {
	MetadataKey string `json:"metadataKey" db:"metadata_key"`
}

type ThoughtMetadataResponse struct {
	Lifetime    string    `json:"lifetime" db:"lifetime"`
	IsBurned    bool      `json:"isBurned" db:"is_burned"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	ThoughtKey  string    `json:"thoughtKey" db:"thought_key"`
}

// TODO: To implement Status table, with id, status, and time when status was changed

type Thought struct {
	Id           int       `json:"id" db:"id"`
	Passphrase   string    `json:"passphrase" db:"passphrase" binding:"required"`
	Lifetime     string    `json:"lifetime" db:"lifetime" binding:"required"`
	IsBurned     bool      `json:"isBurned" db:"is_burned"`
	TimeAccessed time.Time `json:"timeAccessed" db:"time_accessed"`
	CreatedDate  time.Time `json:"createdDate" db:"created_date"`
	MetadataKey  string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey   string    `json:"thoughtKey" db:"thought_key"`
}
