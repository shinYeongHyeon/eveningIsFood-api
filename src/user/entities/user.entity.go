package userentities

import "time"

// User entity
type User struct {
	UUID      string `gorm:"primaryKey"`
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time `sql:"DEFAULT:'current_timestamp'"`
}
