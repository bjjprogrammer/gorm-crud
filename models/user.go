package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name"`
	Email     string `gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"email"`
}
