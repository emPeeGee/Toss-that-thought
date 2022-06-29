package entity

import (
	"github.com/google/uuid"
	"time"
)

type ThoughtCreateInput struct {
	Thought    string    `json:"thought" db:"thought" validate:"required"`
	Passphrase string    `json:"passphrase" db:"passphrase" validate:"max=255"`
	Lifetime   time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

type ThoughtCreateResponse struct {
	MetadataKey           uuid.UUID `json:"metadataKey" db:"metadata_key"`
	ThoughtKey            uuid.UUID `json:"thoughtKey" db:"thought_key"`
	AbbreviatedThoughtKey string    `json:"abbreviatedThoughtKey"`
	IsBurned              bool      `json:"isBurned" db:"is_burned"`
	Lifetime              time.Time `json:"lifetime" db:"lifetime" validate:"required"`
}

// These two are similar
type ThoughtPassphraseInput struct {
	Passphrase string `json:"passphrase" validate:"required,max=255"`
}

type UserPassphrase struct {
	Passphrase string
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
	Lifetime              time.Time  `json:"lifetime" db:"lifetime"`
	AbbreviatedThoughtKey string     `json:"abbreviatedThoughtKey" db:"abbreviated_thought_key" gorm:"column:thought_key"`
	IsBurned              bool       `json:"isBurned" db:"is_burned"`
	BurnedDate            *time.Time `json:"burnedDate" db:"burned_date"`
	IsViewed              bool       `json:"isViewed" db:"is_viewed"`
	ViewedDate            *time.Time `json:"viewedDate" db:"viewed_date"`
	CreatedDate           time.Time  `json:"createdDate" db:"created_date"`
}

type Thought struct {
	Id          int        `json:"id" db:"id"`
	Thought     string     `json:"thought" db:"thought" validate:"required" gorm:"notNull"`
	Passphrase  string     `json:"passphrase" db:"passphrase" validate:"required" gorm:"notNull;size:255"`
	Lifetime    time.Time  `json:"lifetime" db:"lifetime" validate:"required" gorm:"notNull"`
	IsBurned    bool       `json:"isBurned" db:"is_burned" gorm:"default:false;notNull"`
	IsViewed    bool       `json:"isViewed" db:"is_viewed" gorm:"default:false;notNull"`
	BurnedDate  *time.Time `json:"burnedDate" db:"burned_date"`
	ViewedDate  *time.Time `json:"viewedDate" db:"viewed_date"`
	CreatedDate time.Time  `json:"createdDate" db:"created_date" gorm:"autoCreateTime;notNull"`
	MetadataKey uuid.UUID  `json:"metadataKey" db:"metadata_key" gorm:"notNull;type:uuid;default:uuid_generate_v4()"`
	ThoughtKey  uuid.UUID  `json:"thoughtKey" db:"thought_key" gorm:"notNull;type:uuid;default:uuid_generate_v4()"`
}
