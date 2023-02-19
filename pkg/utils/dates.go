package utils

import "time"

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func BeginningOfWeek(date time.Time) time.Time {
	weekday := int(date.Weekday())

	return date.AddDate(0, 0, -weekday)
}

func EndOfWeek(date time.Time) time.Time {
	weekday := int(date.Weekday())

	if weekday == 6 {
		weekday = 0
	}

	return date.AddDate(0, 0, weekday)
}

func RangeBetweenDates(start, end time.Time) (dates []time.Time) {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	days := int(end.Sub(start).Hours()/24) + 1 // +1 to include both start and end date
	if days < 0 {
		days = 0
	}
	ch := make(chan []time.Time)
	go func() {
		defer close(ch)
		dates := []time.Time{}

		for ; !start.After(end); start = start.AddDate(0, 0, 1) {
			dates = append(dates, start)
		}

		ch <- dates
	}()
	return <-ch
}
