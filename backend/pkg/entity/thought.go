package entity

import (
	"time"
)

type ThoughtInput struct {
	Thought    string    `json:"thought" db:"thought" validate:"required"`
	Passphrase string    `json:"passphrase" db:"passphrase" validate:"required,max=255"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtCreateResponse struct {
	MetadataKey string `json:"metadataKey" db:"metadata_key"`
}

type ThoughtMetadataResponse struct {
	Lifetime    string    `json:"lifetime" db:"lifetime"`
	IsBurned    bool      `json:"isBurned" db:"is_burned"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	ThoughtKey  string    `json:"thoughtKey" db:"thought_key"`
}

type ThoughtResponse struct {
	Thought string `json:"thought" db:"thought" validate:"required"`
}

// TODO: To implement Status table, with id, status, and time when status was changed
type Thought struct {
	Id           int       `json:"id" db:"id"`
	Thought      string    `json:"thought" db:"thought" validate:"required"`
	Passphrase   string    `json:"passphrase" db:"passphrase" validate:"required"`
	Lifetime     string    `json:"lifetime" db:"lifetime" validate:"required"`
	IsBurned     bool      `json:"isBurned" db:"is_burned"`
	TimeAccessed time.Time `json:"timeAccessed" db:"time_accessed"`
	CreatedDate  time.Time `json:"createdDate" db:"created_date"`
	MetadataKey  string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey   string    `json:"thoughtKey" db:"thought_key"`
}
