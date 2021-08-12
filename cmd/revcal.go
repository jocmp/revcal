package main

import (
	"fmt"
	"time"

	revcal "github.com/jocmp/revcal/pkg"
)

func main() {
	date := revcal.NewDate(time.Now())
	date.Format()
	fmt.Printf("Today is %s celebrating the %s.\n", date.Format(), date.Symbol())
}
