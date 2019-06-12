package document

import (
	"testing"
)

func TestQueryOne(t *testing.T) {

	uptimeHistory, err := UptimeChange{}.QueryOne()
	if err != nil {
		t.Error(err)
	}
	t.Logf("uptime history: %v \n", uptimeHistory)
}
