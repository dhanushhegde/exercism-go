// Package gigasecond provides function to add gigasecond to a give Date
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond function returns Time after adding gigaseconds to it
func AddGigasecond(t time.Time) time.Time {

	return time.Unix(t.Unix()+1000000000, 0)
}
