package models

import "time"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type LogInfo struct {
	Create_on time.Time `json:"created_on"`
	Action    string    `json:"action"`
}
