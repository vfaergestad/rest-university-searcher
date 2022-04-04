//With inspiration from https://stackoverflow.com/questions/37992660/golang-retrieve-application-uptime

package uptime

// File that includes the startTime variable, and the functions acting upon it

import (
	"time"
)

// The start time of the server
var startTime time.Time

// Init Sets the start time to the current time.
func Init() {
	startTime = time.Now()
}

// GetUptimeString returns the uptime of the server as a string
func GetUptimeString() string {
	return time.Since(startTime).Round(time.Second).String()
}
