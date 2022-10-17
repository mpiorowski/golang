package utils

import (
	"log"
	"os"
	"time"
)

type Email struct {
	To       string
	Template string
	Html     []string
}

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func RoundUpTime(t time.Time, roundOn time.Duration) time.Time {
	t = t.Round(roundOn)

	if time.Since(t) >= 0 {
		t = t.Add(roundOn)
	}

	return t
}

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Error: %s environment variable not set.\n", k)
	}
	return v
}
