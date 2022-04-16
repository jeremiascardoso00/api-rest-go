package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	FirstName   int64 `json:"firstName"`
	LastName    int64 `json:"lastName"`
	Address     int64 `json:"address"`
	PhoneNumber int64 `json:"phoneNumber"`
}
