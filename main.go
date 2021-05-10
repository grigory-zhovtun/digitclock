package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func displayDigits() {
	for i := 0; i < 5; i ++ {
		for j := 0; j < 8; j ++ {
			fmt.Printf("%s ", Display[j][i])
		}
		fmt.Println()
	}
}

func splitDigits(num int) (int, int) {
	a := num / 10
	b := num % 10
	return a, b
}

func setDisplayDigits(ind, num int) {
	switch (num) {
	case 0:
		Display[ind] = Zero
	case 1:
		Display[ind] = One
	case 2:
		Display[ind] = Two
	case 3:
		Display[ind] = Three
	case 4:
		Display[ind] = Four
	case 5:
		Display[ind] = Five
	case 6:
		Display[ind] = Six
	case 7:
		Display[ind] = Seven
	case 8:
		Display[ind] = Eight
	case 9:
		Display[ind] = Nine
	}
}

func drawTime(hour, min, sec int) {
	firstS, secondS := splitDigits(sec)
	firstM, secondM := splitDigits(min)
	firstH, secondH := splitDigits(hour)

	setDisplayDigits(6, firstS)
	setDisplayDigits(7, secondS)

	setDisplayDigits(3, firstM)
	setDisplayDigits(4, secondM)

	setDisplayDigits(0, firstH)
	setDisplayDigits(1, secondH)
}

func main() {
	screen.Clear()

	count := 1 //for separators blink

	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Clock()
		drawTime(hour, min, sec)
		displayDigits()
		time.Sleep(time.Second)

		// animate colons
		count ++ 
		if count == 2 {
			count = 0
			Display[2] = Empty
			Display[5] = Empty
		} else {
			Display[2] = Colon
			Display[5] = Colon
		}
	}
}