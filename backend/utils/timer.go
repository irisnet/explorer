package utils

import "time"

func RunTimer(num int, uint Unit, fn func()) {
	go func() {
		// run once right now
		fn()
		for {
			now := time.Now()
			next := now.Add(ParseDuration(num, uint))
			next = TruncateTime(next, uint)
			t := time.NewTimer(next.Sub(now))
			select {
			case <-t.C:
				fn()
			}
		}
	}()
}
