package main

import (
	"fmt"
	"time"
	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string
	var (
		zero = placeholder{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		}

		one = placeholder{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		}

		two = placeholder{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		}

		three = placeholder{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		}

		four = placeholder{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		}

		five = placeholder{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		}

		six = placeholder{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		}

		seven = placeholder{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		}

		eight = placeholder{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		}

		nine = placeholder{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		}
		colon = placeholder{
			"   ",
			" ░ ",
			"   ",
			" ░ ",
			"   ",
		}
	)
	digits := [...]placeholder{ // length is 10, and it as five string(multi-dimensional array)
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}
	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "   ")
		}
		fmt.Println()
	}
	screen.Clear()
	for {
		screen.MoveTopLeft()
		currentTime := time.Now()
		currentHour := currentTime.Hour()
		currentMinute := currentTime.Minute()
		currentSecond := currentTime.Second()
		clock := [...]placeholder{
			digits[currentHour/10], digits[currentHour%10],
			colon,
			digits[currentMinute/10], digits[currentMinute%10],
			colon,
			digits[currentSecond/10], digits[currentSecond%10],
		}
		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "   ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
