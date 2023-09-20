package model

import "time"

type User struct {
	name     string
	age      int
	birthday time.Time
}
