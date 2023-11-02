package cron

import "time"

func RunByFrequency(f func(), freq time.Duration) {
	f()
	ticker := time.NewTicker(freq)
	for range ticker.C {
		f()
	}
}
