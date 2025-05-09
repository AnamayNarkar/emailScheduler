package utils

import (
	"time"
)

// GetISTLocation returns the Indian Standard Time location
func GetISTLocation() *time.Location {
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Fallback to manual UTC+5:30 if the timezone is not available
		ist = time.FixedZone("IST", 5*60*60+30*60)
	}
	return ist
}

// ConvertToIST converts a time.Time to Indian Standard Time
func ConvertToIST(t time.Time) time.Time {
	return t.In(GetISTLocation())
}

// NowInIST returns the current time in Indian Standard Time
func NowInIST() time.Time {
	return time.Now().In(GetISTLocation())
}
