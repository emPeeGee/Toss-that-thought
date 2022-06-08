package entity

import "database/sql"

type Thought struct {
	Id           int            `db:"id"`
	Passphrase   string         `db:"passphrase"`
	Lifetime     string         `db:"lifetime"`
	CreatedDate  string         `db:"created_date"`
	IsBurned     bool           `db:"is_burned"`
	TimeAccessed sql.NullString `db:"time_accessed"`
	MetadataKey  string         `db:"metadata_key"`
	ThoughtKey   string         `db:"thought_key"`
}
