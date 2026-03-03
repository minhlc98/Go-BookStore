package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID 					string 					`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name 				string 					`gorm:"type:varchar(100);not null" json:"name"`
	Bio 				string 					`gorm:"type:text" json:"bio"`
	CreatedAt 	time.Time 			`gorm:"type:timestamptz(3)" json:"created_at"`
	UpdatedAt 	time.Time 			`gorm:"type:timestamptz(3)" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"type:timestamptz(3);index" json:"deleted_at"`
}	
