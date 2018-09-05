package task

import "testing"

func TestTask(t *testing.T) {
	t.Run("uptime change test", func(t *testing.T) {
		for true {
			UptimeChange()
		}
	})
}
