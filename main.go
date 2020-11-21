package main

import (
	"fmt"
	"time"
)

const dateFormat = "02-01-2006"

func swapDates(dateStart *string, dateEnd *string) {
	// parsed dereference datestart pointer to a time format
	parsedDateStart, err := time.Parse(dateFormat, *dateStart)
	if err != nil {
		println(err)
	}
	start := parsedDateStart.AddDate(0, 0, 7*2).Format(dateFormat) // Added 2 weeks to datestart
	end := parsedDateStart.AddDate(0, 0, 7*3).Format(dateFormat)   // Added 3 weeks to dateStart
	*dateStart = start
	*dateEnd = end
}

func main() {
	dateStart := time.Now().Format(dateFormat) // Get current date
	dateEnd := "None"

	fmt.Printf("DateStart before swap: %s \n", dateStart)
	fmt.Printf("DateEnd before swap: %s \n", dateEnd)

	swapDates(&dateStart, &dateEnd)

	fmt.Printf("DateStart after swap: %s \n", dateStart)
	fmt.Printf("DateEnd after swap: %s \n", dateEnd)

}
