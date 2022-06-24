package entity

import (
	"github.com/lib/pq"
	"time"
)

type ThoughtCreateInput struct {
	Thought    string    `json:"thought" db:"thought" validate:"required"`
	Passphrase string    `json:"passphrase" db:"passphrase" validate:"max=255"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtCreateResponse struct {
	MetadataKey           string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey            string    `json:"thoughtKey" db:"thought_key"`
	AbbreviatedThoughtKey string    `json:"abbreviatedThoughtKey"`
	IsBurned              bool      `json:"isBurned" db:"is_burned"`
	Lifetime              time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtPassphraseInput struct {
	Passphrase string `json:"passphrase" validate:"required,max=255"`
}

type ThoughtResponse struct {
	Thought string `json:"thought" db:"thought"`
}

type ThoughtValidityInformation struct {
	ThoughtKey string    `json:"thoughtKey" db:"thought_key"`
	IsBurned   bool      `json:"isBurned" db:"is_burned"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime"`
	IsViewed   bool      `json:"isViewed" db:"is_viewed"`
}

type ThoughtMetadataResponse struct {
	Lifetime              time.Time   `json:"lifetime" db:"lifetime"`
	AbbreviatedThoughtKey string      `json:"abbreviatedThoughtKey" db:"abbreviated_thought_key"`
	IsBurned              bool        `json:"isBurned" db:"is_burned"`
	BurnedDate            pq.NullTime `json:"burnedDate" db:"burned_date"`
	IsViewed              bool        `json:"isViewed" db:"is_viewed"`
	ViewedDate            pq.NullTime `json:"viewedDate" db:"viewed_date"`
	CreatedDate           time.Time   `json:"createdDate" db:"created_date"`
}

type Thought struct {
	Id          int       `json:"id" db:"id"`
	Thought     string    `json:"thought" db:"thought" validate:"required"`
	Passphrase  string    `json:"passphrase" db:"passphrase" validate:"required"`
	Lifetime    time.Time `json:"lifetime" db:"lifetime" validate:"required"`
	IsBurned    bool      `json:"isBurned" db:"is_burned"`
	IsViewed    bool      `json:"isViewed" db:"is_viewed"`
	BurnedDate  time.Time `json:"burnedDate" db:"burned_date"`
	ViewedDate  time.Time `json:"viewedDate" db:"viewed_date"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	MetadataKey string    `json:"metadataKey" db:"metadata_key"`
	ThoughtKey  string    `json:"thoughtKey" db:"thought_key"`
}
