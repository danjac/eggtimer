package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	invalidInput = "Invalid input, must be H:M:S"
	defaultMsg   = "Time's up"
)

func getInt(s string) (int64, error) {
	value, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return value, err
	}
	return value, nil
}

func main() {

	var (
		value string
		msg   string
	)

	flag.StringVar(&value, "t", "", "H:M:S")
	flag.StringVar(&msg, "m", defaultMsg, "notification message")

	flag.Parse()

	values := strings.Split(value, ":")
	if len(values) != 3 {
		fmt.Println(invalidInput)
		return
	}

	hours, err := getInt(values[0])
	minutes, err := getInt(values[1])
	seconds, err := getInt(values[2])
	if err != nil {
		fmt.Println(invalidInput)
		return
	}

	seconds += (minutes * 60) + (hours * 60 * 60)
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Time left: %d seconds\n", seconds)
		seconds -= 1
		if seconds <= 0 {
			cmd := exec.Command("notify-send", msg)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(msg)
			}
			return
		}
	}

}
