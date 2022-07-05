package thought

import (
	"github.com/google/uuid"
	"time"
)

type CreateDTO struct {
	Thought    string    `json:"thought" validate:"required"`
	Passphrase string    `json:"passphrase" validate:"max=255"`
	Lifetime   time.Time `json:"lifetime" validate:"required"`
}

type CreateResponse struct {
	MetadataKey           uuid.UUID `json:"metadataKey" `
	ThoughtKey            uuid.UUID `json:"thoughtKey"`
	AbbreviatedThoughtKey string    `json:"abbreviatedThoughtKey"`
	IsBurned              bool      `json:"isBurned"`
	Lifetime              time.Time `json:"lifetime" validate:"required"`
}

type PassphraseDTO struct {
	Passphrase string `json:"passphrase" validate:"required,max=255"`
}

type ViewThoughtResponse struct {
	Thought string `json:"thought"`
}

type ValidityInformation struct {
	ThoughtKey string    `json:"thoughtKey"`
	IsBurned   bool      `json:"isBurned"`
	Lifetime   time.Time `json:"lifetime"`
	IsViewed   bool      `json:"isViewed"`
}

type PassphraseInformationResponse struct {
	CanPassphraseBeSkipped bool `json:"canPassphraseBeSkipped"`
}

type MetadataResponse struct {
	MetadataKey           uuid.UUID  `json:"metadataKey" `
	Lifetime              time.Time  `json:"lifetime"`
	AbbreviatedThoughtKey string     `json:"abbreviatedThoughtKey" gorm:"column:thought_key"`
	IsBurned              bool       `json:"isBurned"`
	BurnedDate            *time.Time `json:"burnedDate"`
	IsViewed              bool       `json:"isViewed"`
	ViewedDate            *time.Time `json:"viewedDate"`
	CreatedDate           time.Time  `json:"createdDate"`
}
