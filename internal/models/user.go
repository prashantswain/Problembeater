package models

import "time"

type Student struct {
	Id             int64      `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
	Name           string     `json:"name"`
	MobileNumber   int        `json:"mobileNumber"`
	Email_Id       string     `json:"emailID"`
	Gender         *string    `json:"gender"`
	Age            *int       `json:"age"`
	Password       *string    `json:"password,omitempty"`
	Class          *Class     `json:"class"`
	ProfilePicture *string    `json:"profilePicture"`
}
