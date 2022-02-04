//With inspiration from https://stackoverflow.com/questions/37992660/golang-retrieve-application-uptime

package uptime

import "time"

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func GetUptime() int {
	return int(time.Since(startTime).Seconds())
}
