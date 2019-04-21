package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var deadlineStr string
	flag.StringVar(&deadlineStr, "deadline", "", "deadline date(yyyy-mm-dd)")
	flag.Parse()

	if err := checkDeadline(deadlineStr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	today := time.Now()
	const layout = "2006-01-02"
	deadline, _ := time.Parse(layout, deadlineStr)

	duration := deadline.Sub(today)
	days := int(duration.Hours()) / 24

	if days > 30 {
		fmt.Printf("Remain : \x1b[32m%ddays\n", days)
	} else if days > 10 {
		fmt.Printf("Remain : \x1b[33m%ddays\n", days)
	} else {
		fmt.Printf("Remain : \x1b[31m%ddays\n", days)
	}
}

func checkDeadline(deadlineStr string) error {
	if deadlineStr == "" {
		msg := "Usage ./daysuntil -deadline=\"yyyy-mm-dd\""
		return fmt.Errorf("err %s", msg)
	}
	return nil
}
