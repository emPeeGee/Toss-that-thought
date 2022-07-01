package entity

import (
	"github.com/google/uuid"
	"time"
)

type Thought struct {
	Id          int        `json:"id"`
	Thought     string     `json:"thought" validate:"required" gorm:"notNull"`
	Passphrase  string     `json:"passphrase" validate:"required" gorm:"notNull;size:255"`
	Lifetime    time.Time  `json:"lifetime" validate:"required" gorm:"notNull"`
	IsBurned    bool       `json:"isBurned" gorm:"default:false;notNull"`
	IsViewed    bool       `json:"isViewed" gorm:"default:false;notNull"`
	BurnedDate  *time.Time `json:"burnedDate"`
	ViewedDate  *time.Time `json:"viewedDate"`
	CreatedDate time.Time  `json:"createdDate" gorm:"autoCreateTime;notNull"`
	MetadataKey uuid.UUID  `json:"metadataKey" gorm:"notNull;type:uuid;default:uuid_generate_v4()"`
	ThoughtKey  uuid.UUID  `json:"thoughtKey" gorm:"notNull;type:uuid;default:uuid_generate_v4()"`
}
