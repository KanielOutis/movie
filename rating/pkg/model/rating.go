package model

type RecordID string
type RecordType string

const (
	RecordTypeMovie RecordType = "movie"
)

type UserID string
type RatingValue int

// Rating represents a rating for a records.
type Rating struct {
	RecordID   RecordID    `json:"record_id"`
	UserID     UserID      `json:"user_id"`
	Value      RatingValue `json:"value"`
	RecordType RecordType  `json:"record_type"`
}
