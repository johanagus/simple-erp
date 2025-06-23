package domain

import "time"

type Store struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Address   string    `gorm:"type:varchar(255)" json:"address" validate:"required"`
	CreatedAt time.Time `json:"-" grom:"<-:create"`
	UpdatedAt time.Time `json:"-" gorm:"<-:update"`
	DeletedAt time.Time `json:"-" gorm:"<-:delete"`
}
