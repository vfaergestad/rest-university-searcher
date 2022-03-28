package uptime

import (
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	testTime = time.Now()
	Init()
	m.Run()
}

func TestGetUptimeString(t *testing.T) {
	expected := time.Since(testTime).Round(time.Second).String()
	actual := GetUptimeString()
	if actual == "" {
		t.Error("Uptime string is empty")
	}
	if actual != expected {
		t.Errorf("Uptime string is not correct. Expected: %s, Actual: %s", expected, actual)
	}
}
