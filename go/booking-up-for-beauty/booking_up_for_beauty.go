package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	//fmt.Printf("\nSchedule() Input Date: %s\n", date)
	// Format mask for output
	//layout := "January 02, 2006 15:04:05"
	layout := "1/2/2006 15:04:05"
	parsed, err := time.Parse(layout, date)
	if err != nil {
		layout2 := "January 2, 2006 15:04:05"
		parsed2, err2 := time.Parse(layout2, date)
		if err2 != nil {
			layout3 := "Monday, January 2, 2006 15:04:05"
			parsed3, _ := time.Parse(layout3, date)
			//fmt.Printf("\nParsed3: %s\n", parsed3)
			return parsed3
		}
		//fmt.Printf("\nParsed2: %s\n", parsed2)
		return parsed2
	}
	//fmt.Printf("\nParsed: %s\n", parsed)
	return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	appointmentDate := Schedule((date))
	return now.After(appointmentDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	//fmt.Printf("\nIs AfternooApointmetnt() Input Date is %s", date)
	parsed := Schedule(date)
	hour := parsed.Hour()
	return hour >= 12 && hour <= 18

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	//fmt.Printf("\nDescription() Input is: %s\n", date)
	parsed := Schedule(date)
	//fmt.Printf("\nparsed description() input is: %s\n", parsed)
	dow := parsed.Weekday()
	year, _, day := parsed.Date()
	month := parsed.Month()
	hour := parsed.Hour()
	minute := parsed.Minute()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dow, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	year, _, _ := now.Date()
	return time.Date(year, 9, 15, 0, 0, 0, 0, time.UTC)
}
