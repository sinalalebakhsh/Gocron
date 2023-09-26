package features

var TitleOfTimeData = []string{
    "ALL TIME",
    "ALLTIME",
}

type DataBase struct{ Alldatafield string }

var OriginalTimeData = DataBase{
	Alldatafield: `
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
████████████████████████████████████████████████████████████████████████
255.The Methods for Working with Time Values
    Name                Description
    ----------------    -------------------------------------------------
    Add(duration)       This method adds the specified Duration to the Time and returns the result.
    Sub(time)           This method returns a Duration that expresses the difference between the Time on
                        which the method has been called and the Time provided as the argument.
    AddDate(y, m, d)    This method adds the specified number of years, months, and days to the Time and
                        returns the result.
    After(time)         This method returns true if the Time on which the method has been called occurs
                        after the Time provided as the argument.
    Before(time)        This method returns true if the Time on which the method has been called occurs
                        before the Time provided as the argument.
    Equal(time)         This method returns true if the Time on which the method has been called is equal
                        to the Time provided as the argument.
    IsZero()            This method returns true if the Time on which the method has been called
                        represents the zero-time instant, which is January 1, year 1, 00:00:00 UTC.
    In(loc)             This method returns the Time value, expressed in the specified Location.
    Location()          This method returns the Location that is associated with the Time, effectively
                        allowing a time to be expressed in a different time zone.
    Round(duration)     This method rounds the Time to the nearest interval represented by a Duration
                        value.
    Truncate(duration)  This method rounds the Time down to the nearest interval represented by a
                        Duration value.
████████████████████████████████████████████████████████████████████████
256.time.Parse()
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
            if err == nil {
                Printfln("After: %v", t.After(time.Now()))
                Printfln("Round: %v", t.Round(time.Hour))
                Printfln("Truncate: %v", t.Truncate(time.Hour))
            } else {
                fmt.Println(err.Error())
            }
        }
    Output:
        After: false
        Round: 1995-06-09 05:00:00 +0100 BST
        Truncate: 1995-06-09 04:00:00 +0100 BST
████████████████████████████████████████████████████████████████████████
257.Comparing Time Values
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
            t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
            
            Printfln("Equal Method: %v", t1.Equal(t2))
            Printfln("Equality Operator: %v", t1 == t2)
        }
    Output:
        Equal Method: true
        Equality Operator: false
████████████████████████████████████████████████████████████████████████
258.The Duration Constants in the time Package
    Name            Description
    ------------    ----------------------------------------
    Hour            This constant represents 1 hour.
    Minute          This constant represents 1 minute.
    Second          This constant represents 1 second.
    Millisecond     This constant represents 1 millisecond.
    Microsecond     This constant represents 1 microsecond.
    Nanosecond      This constant represents 1 nanosecond.
████████████████████████████████████████████████████████████████████████
259.The Duration Methods
    Name                Description
    ----------------    ---------------------------------------------
    Hours()             This method returns a float64 that represents the Duration in hours.
    Minutes()           This method returns a float64 that represents the Duration in minutes.
    Seconds()           This method returns a float64 that represents the Duration in seconds.
    Milliseconds()      This method returns an int64 that represents the Duration in milliseconds.
    Microseconds()      This method returns an int64 that represents the Duration in microseconds.
    Nanoseconds()       This method returns an int64 that represents the Duration in nanoseconds.
    Round(duration)     This method returns a Duration, which is rounded to the nearest multiple of the
                        specified Duration.
    Truncate(duration)  This method returns a Duration, which is rounded down to the nearest multiple of
                        the specified Duration.
████████████████████████████████████████████████████████████████████████
260.Hours() - Minutes() - Seconds() - rounded.Hours() - rounded.Minutes()
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            var d time.Duration = time.Hour + (30 * time.Minute)
            Printfln("Hours: %v", d.Hours())
            Printfln("Mins: %v", d.Minutes())
            Printfln("Seconds: %v", d.Seconds())
            Printfln("Millseconds: %v", d.Milliseconds())
        
        
            rounded := d.Round(time.Hour)
            Printfln("Rounded Hours: %v", rounded.Hours())
            Printfln("Rounded Mins: %v", rounded.Minutes())
        
            trunc := d.Truncate(time.Hour)
            Printfln("Truncated Hours: %v", trunc.Hours())
            Printfln("Rounded Mins: %v", trunc.Minutes())
        }
    Output:
        Hours: 1.5
        Mins: 90
        Seconds: 5400
        Millseconds: 5400000
        Rounded Hours: 2
        Rounded Mins: 120
        Truncated Hours: 1
        Rounded Mins: 60
████████████████████████████████████████████████████████████████████████
261.The time Functions for Creating Duration Values relative to a Time
    Name            Description
    -----------     ----------------------------------------
    Since(time)     This function returns a Duration expressing the elapsed time since the specified Time value.
    Until(time)     This function returns a Duration expressing the elapsed time until the specified Time value.
████████████████████████████████████████████████████████████████████████
262.time.Until(time) - Since(time)
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            toYears := func(d time.Duration) int {
                return int(d.Hours() / (24 * 365))
            }
            
            future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
            past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
            
            Printfln("this year is %v.",time.Now().Year())
            Printfln("Future: %v is %v.", toYears(time.Until(future)), future.Year())
            Printfln("Past: %v is %v.", toYears(time.Since(past)), past.Year())
        }
    Output:
        this year is 2023.
        Future: 27 is 2050.
        Past: 58 is 1964.
████████████████████████████████████████████████████████████████████████
263.time.ParseDuration function
    This function returns a Duration and an error, indicating if there were problems
    parsing the specified string.
    The format of the strings supported by the ParseDuration function is a sequence of number values
    followed by the unit indicators:
    Unit        Description
    -----       --------------------
    h           This unit denotes hours.
    m           This unit denotes minutes.
    s           This unit denotes seconds.
    ms          This unit denotes milliseconds.
    us or μs    These units denotes microseconds.
    ns          This unit denotes nanoseconds.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func main() {
            d, err := time.ParseDuration("1h30m")
            if err == nil {
                Printfln("Hours: %v", d.Hours())
                Printfln("Mins: %v", d.Minutes())
                Printfln("Seconds: %v", d.Seconds())
                Printfln("Millseconds: %v", d.Milliseconds())
            } else {
                fmt.Println(err.Error())
            }
        }
    Output:
        Hours: 1.5
        Mins: 90
        Seconds: 5400
        Millseconds: 5400000
████████████████████████████████████████████████████████████████████████
264.Using the Time Features for Goroutines and Channels
    
    The time package provides a small set of functions that are useful for working with goroutines and channels.

    The time Package Functions:
    Name                        Description
    ----------------------      ----------------------------------------
    Sleep(duration)             This function pauses the current goroutine for at least the specified duration.
    AfterFunc(duration,func)    This function executes the specified function in its own goroutine after the
                                specified duration. The result is a *Timer, whose Stop method can be used to
                                cancel the execution of the function before the duration elapses.
    After(duration)             This function returns a channel that blocks for the specified duration and then
                                yields a Time value. See the “Receiving Timed Notifications” section for details.
    Tick(duration)              This function returns a channel that periodically sends a Time value, where the
                                period is specified as a duration.
████████████████████████████████████████████████████████████████████████
265.time.Sleep(time.Second * 1)
    Pausing a Goroutine
    The Sleep function pauses execution of the current goroutine for a specified duration

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            
            go writeToChannel(nameChannel)
            
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
    Read name: Alice
                        // 1 second delaying
    Read name: Bob
                        // 1 second delaying
    Read name: Charlie
                        // 1 second delaying
    Read name: Dora
████████████████████████████████████████████████████████████████████████
266.time.AfterFunc() function
    The AfterFunc function is used to defer the execution of a function for a specified period
    Deferring a Function:
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                // time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
        
            time.AfterFunc(time.Second*5, func() {
                writeToChannel(nameChannel)
            })
        
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
    // It waits for 5 seconds and then continues the program.
    Read name: Alice
    Read name: Bob
    Read name: Charlie
    Read name: Dora
████████████████████████████████████████████████████████████████████████
267.time.After(time.Second * 2)
    The result from the After function is a channel that carries Time values. The channel blocks for the
    specified duration, when a Time value is sent, indicating the duration has passed. In this example, the
    value sent over the channel acts as a signal and is not used directly, which is why it is assigned to the blank
    identifier, like this:

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
        
            Printfln("Waiting for initial duration...")
            _ = <-time.After(time.Second * 2)
            Printfln("Initial duration elapsed.")
        
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 1)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
        Waiting for initial duration... // Wait for 2 seconds
        Initial duration elapsed.
        Read name: Alice    // wait for 1 second
        Read name: Bob      // wait for 1 second
        Read name: Charlie  // wait for 1 second
        Read name: Dora     // wait for 1 second
████████████████████████████████████████████████████████████████████████
268.time.Sleep(time.Second * 3) with select statement
    Using a Timeout in a Select Statement

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            Printfln("Waiting for initial duration...")
            _ = <-time.After(time.Second * 2)
            Printfln("Initial duration elapsed.")
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
                time.Sleep(time.Second * 3)
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            channelOpen := true
            for channelOpen {
                Printfln("Starting channel read")
                select {
                case name, ok := <-nameChannel:
                    if !ok {
                        channelOpen = false
                    } else {
                        Printfln("Read name: %v", name)
                    }
                case <-time.After(time.Second * 2):
                    Printfln("Timeout")
                }
            }
        }
    Output:
        Starting channel read
        Waiting for initial duration...
        Timeout
        Starting channel read
        Initial duration elapsed.
        Read name: Alice
        Starting channel read
        Timeout
        Starting channel read
        Read name: Bob
        Starting channel read
        Timeout
        Starting channel read
        Read name: Charlie
        Starting channel read
        Timeout
        Starting channel read
        Read name: Dora
        Starting channel read
        Timeout
        Starting channel read
████████████████████████████████████████████████████████████████████████
269.NewTimer(duration)
    This function returns a *Timer with the specified period.
    Caution Be careful when stopping a timer. 
    The timer's channel is not closed, which means that reads from
    the channel will continue to block even after the timer has stopped.

    The Methods Defined by the Timer Struct:
    Name                Description
    ------------        -------------------------------------------
    C                   This field returns the channel over which the Time will send its Time value.
    Stop()              This method stops the timer. The result is a bool that will be true if the timer has been
                        stopped and false if the timer had already sent its message.
    Reset(duration)     This method stops a timer and resets it so that its interval is the specified Duration.
    
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(channel chan<- string) {
            timer := time.NewTimer(time.Minute * 10)
            go func() {
                time.Sleep(time.Second * 2)
                Printfln("Resetting timer")
                timer.Reset(time.Second)
            }()
            Printfln("Waiting for initial duration...")
            <-timer.C
            Printfln("Initial duration elapsed.")
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            for _, name := range names {
                channel <- name
            }
            close(channel)
        }
        func main() {
            nameChannel := make(chan string)
            go writeToChannel(nameChannel)
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output:
        Waiting for initial duration...
        Resetting timer
        Initial duration elapsed.
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
████████████████████████████████████████████████████████████████████████
270.time.Tick(time.Second)
    Receiving Recurring Notifications دریافت اعلان های مکرر
    The Tick function returns a channel over which Time values are sent at a specified interval
    
    Tick is a convenience wrapper for NewTicker providing access to the ticking channel only. 
    While Tick is useful for clients that have no need to shut down the Ticker, 
    be aware that without a way to shut it down the underlying
    Ticker cannot be recovered by the garbage collector; it "leaks".
    Unlike NewTicker, Tick will return nil if d <= 0.

    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(nameChannel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            tickChannel := time.Tick(time.Second)
            index := 0
            for {
                <-tickChannel
                nameChannel <- names[index]
                index++
                if index == len(names) {
                    index = 0
                }
            }
        }
        func main() {
            nameChannel := make(chan string)

            go writeToChannel(nameChannel)

            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
            
        }
    Output:
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora
        ...
            
████████████████████████████████████████████████████████████████████████
271.NewTicker(duration)
    The result of the NewTicker function is a pointer to a Ticker struct, which defines the field and methods

    The time Function for Creating a Ticker:
    Name                    Description
    ---------------------   ---------------------------------
    NewTicker(duration)     This function returns a *Ticker with the specified period.

    The Field and Methods Defined by the Ticker Struct:
    Name                Description
    ----------------    --------------------------------
    C                   This field returns the channel over which the Ticker will send its Time values.
    Stop()              This method stops the ticker (but does not close the channel returned by the C field).
    Reset(duration)     This method stops a ticker and resets it so that its interval is the specified Duration.
████████████████████████████████████████████████████████████████████████
272.Creating a Ticker in the main.go
    example:
        package main
        import (
            "fmt"
            "time"
        )
        func Printfln(template string, values ...interface{}) {
            fmt.Printf(template+"\n", values...)
        }
        func writeToChannel(nameChannel chan<- string) {
            names := []string{"Alice", "Bob", "Charlie", "Dora"}
            ticker := time.NewTicker(time.Second / 10)
            index := 0
            for {
                <-ticker.C
                nameChannel <- names[index]
                index++
                if index == len(names) {
                    ticker.Stop()
                    close(nameChannel)
                    break
                }
            }
        }
        func main() {
            nameChannel := make(chan string)
        
            go writeToChannel(nameChannel)
        
            for name := range nameChannel {
                Printfln("Read name: %v", name)
            }
        }
    Output: 
        // This is printed after milliseconds
        Read name: Alice
        Read name: Bob
        Read name: Charlie
        Read name: Dora`,
}
