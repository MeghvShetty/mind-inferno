package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	time, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	dateConverstion, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}

	return dateConverstion.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {

	dateConverstion, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	if dateConverstion.Hour() >= 12 && dateConverstion.Hour() < 18 {
		return true
	}
	return false

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateConverstion, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("You have an appointment on %s", dateConverstion.Format("Monday, January 2, 2006, at 15:04."))

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

//func main() {
// 	p := fmt.Println
// 	p(HasPassed("July 25, 2023 13:45:00"))
// 	p(IsAfternoonAppointment("Thursday, July 25, 2019 18:45:00"))
// 	p(Description("6/6/2005 10:30:00"))
// 	p(AnniversaryDate())
// }
//
