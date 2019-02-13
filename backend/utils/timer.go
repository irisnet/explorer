package utils

import "time"

func RunTimer(d time.Duration, prec Prec, fn func()) {
	go func() {
		// run once right now
		fn()
		for {
			now := time.Now()
			next := now.Add(d)
			next = TruncateTime(next, prec)
			t := time.NewTimer(next.Sub(now))
			select {
			case <-t.C:
				fn()
			}
		}
	}()
}
