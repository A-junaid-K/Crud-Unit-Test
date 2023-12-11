package models

type Ussankutti struct {
	ID    uint   `json:"id" gorm:"primaryKey;unique"`
	Name  string `json:"name" gorm:"not null" validate:"min=5,max=20"`
	Email string `json:"email" gorm:"not null" validate:"email"`
	Phone string `json:"phone" gorm:"not null" validate:"min=10,max=10"`
}
