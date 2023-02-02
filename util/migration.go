package util

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp"`
	Name      string    `gorm:"column:name;unique"`
	Password  string    `gorm:"column:password"`
}
