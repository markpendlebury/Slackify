package main

import (
	"time"
)

func GetExpiryDate(expiresIn int) time.Time {
	expiryDate := time.Now().Local().Add(time.Hour * 1)
	return expiryDate
}
