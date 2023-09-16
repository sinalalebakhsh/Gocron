package features

type TimeData struct {
	alltimeData string
}


var OriginalTimeData TimeData = TimeData{
	alltimeData: `
243.Putting Dates, Times, and Durations in Context
What are they?
The features provided by the time package are used to represent
specific moments in time and intervals or durations.

Why are they useful?
These features are useful in any application that needs to deal with
calendaring or alarm and for the development of any feature that
requires delays or notifications in the future.

How are they used?
The time package defines data types for representing dates and
individual units of time and functions for manipulating them. There
are also features integrated into the Go channel system.

Are there any pitfalls or limitations?
Dates can be complex, and care must be taken to deal with calendar
and time zone issues.

Are there any alternatives?
These are optional features, and their use is not required.

████████████████████████████████████████████████████████████████████████
244.The Functions in the time Package for Creating Time Values
Name                                    Description
--------                                -----------------------------
Now()                                   This function creates a Time representing the current moment in time.

Date(y, m, d, h, min, sec, nsec, loc)   This function creates a Time representing a specified moment in time, which is
										expressed by the year, month, day, hour, minute, second, nanosecond, and Location
										arguments. (The Location type is described in the “Parsing Time Values from Strings”
										section.)

Unix(sec, nsec)                         This function creates a Time value from the number of seconds and nanoseconds since
										January 1, 1970, UTC, commonly known as Unix time.
									
████████████████████████████████████████████████████████████████████████
245.The Methods for Accessing Time Components
Name            Description
--------        ----------------------------
Date()          This method returns the year, month, and day components. The year and day are
				expressed as int values and the month as a Month value.

Clock()         This method returns the hour, minutes, and seconds components of the Time.
											
Year()          This method returns the year component, expressed as an int.
											
YearDay()       This method returns the day of the year, expressed as an int between 1 and 366 (to accommodate leap years).

Month()         This method returns the month component, expressed using the Month type.

Day()           This method returns the day of the month, expressed as an int.

Weekday()       This method returns the day of the week, expressed as a Weekday.

Hour()          This method returns the hour of the day, expressed as an int between 0 and 23.

Minute()        This method returns the number of minutes elapsed into the hour of the day, expressed as an int between 0 and 59.

Second()        This method returns the number of seconds elapsed into the minute of the hour, expressed as an int between 0 and 59.
											
Nanosecond()    This method returns the number of nanoseconds elapsed into the second of the minute,
				expressed as an int between 0 and 999,999,999.
████████████████████████████████████████████████████████████████████████
246.The Types Used to Describe Time Components
Name        Description
-------     ------------------------------------------------------------------------
Month       This type represents a month, and the time package defines constant values for the English-
			language month names: January, February, etc. The Month type defines a String method that
			uses these names when formatting strings.

Weekday     This type represents a day of the week, and the time package defines constant values for the
			English-language weekday names: Sunday, Monday, etc. The Weekday type defines a String
			method that uses these names when formatting strings.
████████████████████████████████████████████████████████████████████████
247.Creating Time Values
example:
	package main
	import "time"
	func PrintTime(label string, t *time.Time) {
		Printfln("%s: Day: %v: Month: %v Year: %v",
			label, t.Day(), t.Month(), t.Year())
	}
	func main() {
		current := time.Now()
		specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
		unix := time.Unix(1433228090, 0)
		PrintTime("Current", &current)
		PrintTime("Specific", &specific)
		PrintTime("UNIX", &unix)
	}
Output:
	Current: Day: 15: Month: September Year: 2023
	Specific: Day: 9: Month: June Year: 1995
	UNIX: Day: 2: Month: June Year: 2015

example:
	package main
	import "time"
	func PrintTime(label string, t *time.Time) {
		Printfln("%s: Day: %v: Month: %v Year: %v",
			label, t.Day(), t.Month(), t.Year())
	}
	func main() {
		current := time.Now()
		specific := time.Date(1993, time.June, 0, 0, 0, 0, 0, time.Local)
		unix := time.Unix(0, 0)
		PrintTime("Current", &current)
		PrintTime("Specific", &specific)
		PrintTime("UNIX", &unix)
	}
Output:
	Current: Day: 15: Month: September Year: 2023
	Specific: Day: 31: Month: May Year: 1993
	UNIX: Day: 31: Month: December Year: 1969

████████████████████████████████████████████████████████████████████████
248.The Time Method for Creating Formatted Strings
Name                Description
--------------      --------------------
Format(layout)      This method returns a formatted string, which is created using the specified layout.

example:
	package main
	import (
		"fmt"
		"time"
	)
	func PrintTime(label string, t *time.Time) {
		layout := "Day: 02 Month: Jan Year: 2006"
		fmt.Println(label, t.Format(layout))
	}
	func main() {
		current := time.Now()
		specific := time.Date(1993, time.June, 0, 0, 0, 0, 0, time.Local)
		unix := time.Unix(0, 0)
		PrintTime("Current", &current)
		PrintTime("Specific", &specific)
		PrintTime("UNIX", &unix)
	}
Output:
	Current Day: 16 Month: Sep Year: 2023
	Specific Day: 31 Month: May Year: 1993
	UNIX Day: 31 Month: Dec Year: 1969
████████████████████████████████████████████████████████████████████████
249.The Layout Constants Defined by the time Package
Name            Reference Date Format
-----------     ----------------------------------
ANSIC           Mon Jan _2 15:04:05 2006
UnixDate        Mon Jan _2 15:04:05 MST 2006
RubyDate        Mon Jan 02 15:04:05 -0700 2006
RFC822          02 Jan 06 15:04 MST
RFC822Z         02 Jan 06 15:04 -0700
RFC850          Monday, 02-Jan-06 15:04:05 MST
RFC1123         Mon, 02 Jan 2006 15:04:05 MST
RFC1123Z        Mon, 02 Jan 2006 15:04:05 -0700
RFC3339         2006-01-02T15:04:05Z07:00
RFC3339Nano     2006-01-02T15:04:05.999999999Z07:00
Kitchen         3:04PM
Stamp           Jan _2 15:04:05
StampMilli      Jan _2 15:04:05.000
StampMicro      Jan _2 15:04:05.000000
StampNano       Jan _2 15:04:05.000000000
████████████████████████████████████████████████████████████████████████
250.The time Package Functions for Parsing Strings into Time Values
Name                                        Description
--------------------------------------      -----------------------------------
Parse(layout, str)                          This function parses a string using the specified layout to create a Time value.
											An error is returned to indicate problems parsing the string.

ParseInLocation(layout, str, location)      This function parses a string, using the specified layout and using the
											Location if no time zone is included in the string. An error is returned to
											indicate problems parsing the string.

example:
	package main
	import (
		"fmt"
		"time"
	)
	func PrintTime(label string, t *time.Time) {
		fmt.Println(label, t.Format(time.RFC822Z))
	}
	func main() {
		dates := []string{
			"09 Jun 95 00:00 GMT",
			"02 Jun 15 00:00 GMT",
		}
		
		for _, d := range dates {
			time, err := time.Parse(time.RFC822, d)
			if err == nil {
				PrintTime("Parsed", &time)
			} else {
				Printfln("Error: %s", err.Error())
			}
		}
	}
Output:
	Parsed 09 Jun 95 00:00 +0000
	Parsed 02 Jun 15 00:00 +0000
████████████████████████████████████████████████████████████████████████
251.time.ParseInLocation()
example:
	package main
	import (
		"fmt"
		"time"
	)
	func PrintTime(label string, t *time.Time) {
		//layout := "Day: 02 Month: Jan Year: 2006"
		fmt.Println(label, t.Format(time.RFC822Z))
	}
	func main() {
		layout := "02 Jan 06 15:04"
		date := "09 Jun 95 19:30"
		london, lonerr := time.LoadLocation("Europe/London")
		newyork, nycerr := time.LoadLocation("America/New_York")
		Tehran, TehranErr := time.LoadLocation("Asia/Tehran")
	
		if lonerr == nil && nycerr == nil  && TehranErr == nil{
	
			TehranTime, _ := time.ParseInLocation(layout, date, Tehran)
			PrintTime("Tehran:",&TehranTime)
	
			nolocation, _ := time.Parse(layout, date)
			PrintTime("No location:", &nolocation)
	
			londonTime, _ := time.ParseInLocation(layout, date, london)
			PrintTime("London:", &londonTime)
	
			newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
			PrintTime("New York:", &newyorkTime)
	
		} else {
			fmt.Println(lonerr.Error(), nycerr.Error())
		}
	}
Output:
	Tehran: 09 Jun 95 19:30 +0430
	No location: 09 Jun 95 19:30 +0000
	London: 09 Jun 95 19:30 +0100
	New York: 09 Jun 95 19:30 -0400
████████████████████████████████████████████████████████████████████████
252.The Functions for Creating Locations
Name                                Description
-------------------------           -----------------------------------------
LoadLocation(name)                  This function returns a *Location for the specified name and an
									error that indicates any problems.

									LoadLocationFromTZData(name,data)   This function returns a *Location from a byte slice that contains a
									formatted time zone database.

FixedZone(name, offset)             This function returns a *Location that always uses the specified
									name and offset from UTC.

The place names are defined in the IANA time zone database, 
https://www.iana.org/time-zones , 
and are listed by 
https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

████████████████████████████████████████████████████████████████████████
253.Embedding The Time Zone Database
example:
	package main
	import (
		"fmt"
		"time"
	)
	func PrintTime(label string, t *time.Time) {
		//layout := "Day: 02 Month: Jan Year: 2006"
		fmt.Println(label, t.Format(time.RFC822Z))
	}
	func main() {
		layout := "02 Jan 06 15:04"
		date := "09 Jun 95 19:30"
		Tehran, TehranErr := time.LoadLocation("Asia/Tehran")
		local, _ := time.LoadLocation("Local")
		if TehranErr == nil{
			TehranTime, _ := time.ParseInLocation(layout, date, Tehran)
			PrintTime("Tehran:",&TehranTime)
			localTime, _ := time.ParseInLocation(layout, date, local)
			PrintTime("Local:", &localTime)
			nolocation, _ := time.Parse(layout, date)
			PrintTime("No location:", &nolocation)
		} else {
			fmt.Println(TehranErr.Error())
		}
	}
Output:
	Tehran: 09 Jun 95 19:30 +0430
	Local: 09 Jun 95 19:30 -0400
	No location: 09 Jun 95 19:30 +0000
████████████████████████████████████████████████████████████████████████
254.Specifying Time Zones
The arguments to the FixedZone function are a name and the number of seconds offset from UTC. This
example creates three fixed time zones, one of which is an hour ahead of UTC, one of which is four hours
behind, and one of which has no offset.
example:
	package main
	import (
		"fmt"
		"time"
	)
	func PrintTime(label string, t *time.Time) {
		//layout := "Day: 02 Month: Jan Year: 2006"
		fmt.Println(label, t.Format(time.RFC822Z))
	}
	func main() {
		layout := "02 Jan 06 15:04"
		date := "09 Jun 95 19:30"
		london := time.FixedZone("BST", 1*60*60)
		newyork := time.FixedZone("EDT", -4*60*60)
		local := time.FixedZone("Local", 0)
		
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)
		
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
	}
Output:
	No location: 09 Jun 95 19:30 +0000
	London: 09 Jun 95 19:30 +0100
	New York: 09 Jun 95 19:30 -0400
	Local: 09 Jun 95 19:30 +0000
████████████████████████████████████████████████████████████████████████`,
} 
