package user

import "time"

type Student struct {
	Id           int64
	CreatedAt    time.Time `json:"createdAt"`
	Name         string    `json:"name"`
	MobileNumber string    `json:"mobileNumber"`
	Email_Id     string    `json:"emailID"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Password     string    `json:"password,omitempty"`
}
