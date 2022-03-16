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

// GetUptimeInSeconds returns the time in seconds since the last restart
func GetUptimeInSeconds() int {
	return int(time.Since(startTime).Seconds())
}

func GetUptimeString() string {
	return time.Since(startTime).String()
}
