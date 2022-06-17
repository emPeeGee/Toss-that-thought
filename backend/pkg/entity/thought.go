package entity

import (
	"time"
)

// Todo: Rename with Create suffix
type ThoughtCreateInput struct {
	Thought    string    `json:"thought" db:"thought" validate:"required"`
	Passphrase string    `json:"passphrase" db:"passphrase" validate:"max=255"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtPassphraseInput struct {
	Passphrase string `json:"passphrase" validate:"required,max=255"`
}

type ThoughtResponse struct {
	Thought string `json:"thought" db:"thought"`
}

type ThoughtCreateResponse struct {
	MetadataKey string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey  string    `json:"thoughtKey" db:"thought_key"`
	IsBurned    bool      `json:"isBurned" db:"is_burned"`
	Lifetime    time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtMetadataResponse struct {
	Lifetime              string    `json:"lifetime" db:"lifetime"`
	AbbreviatedThoughtKey string    `json:"abbreviatedThoughtKey" db:"abbreviated_thought_key"`
	IsBurned              bool      `json:"isBurned" db:"is_burned"`
	CreatedDate           time.Time `json:"createdDate" db:"created_date"`
}

// TODO: To implement Status table, with id, status, and time when status was changed
// TODO: Passphrase should be encrypted
type Thought struct {
	Id          int       `json:"id" db:"id"`
	Thought     string    `json:"thought" db:"thought" validate:"required"`
	Passphrase  string    `json:"passphrase" db:"passphrase" validate:"required"`
	Lifetime    string    `json:"lifetime" db:"lifetime" validate:"required"`
	IsBurned    bool      `json:"isBurned" db:"is_burned"`
	TimeBurned  time.Time `json:"timeBurned" db:"time_burned"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	MetadataKey string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey  string    `json:"thoughtKey" db:"thought_key"`
}
