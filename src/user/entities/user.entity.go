package user_entities

import "time"

type User struct {
	ID 	 	  int       `gorm:"primaryKey"`
	Name 	  string
	Password  string
	CreatedAt time.Time `sql:"DEFAULT:'current_timestamp'"`
}

