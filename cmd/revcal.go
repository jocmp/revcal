package main

import (
	"fmt"
	"os"
	"time"

	revcal "github.com/jocmp/revcal/pkg"
)

func main() {
	time, err := findTime(os.Args)
	if err != nil {
		fmt.Println("Couldn't find the time")
	}
	date := revcal.NewDate(time)
	date.Format()
	fmt.Printf("Today is %s celebrating the %s.\n", date.Format(), date.Symbol())
}

func findTime(args []string) (time.Time, error) {
	if len(args) > 1 {
		argTime, err := time.Parse("2006-01-02", os.Args[1])
		if err != nil {
			return argTime, err
		}
		return argTime, nil
	} else {
		return time.Now(), nil
	}
}
