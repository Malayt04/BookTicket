package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const(
    Manager UserRole = "manager"
    Attendee UserRole = "attendee"
)

type User struct{
    ID uint `json:"id" gorm:"primarykey"`
    Name string `json:"name"`
    Email string `json:"email"`
    Role UserRole `json:"role"`
    Password string `json:"password"`
    Tickets []Ticket `json:"tickets"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


func (u *User) AfterCreate(db *gorm.DB)(err error){
    if u.ID == 1{
        db.Model(u).Update("role", Manager)
    }

    return
}