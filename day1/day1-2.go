package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	increasecount := -1
	var oldsum int64 = 0
	var newsum int64 = 0
	window := [3]int64{0, 0, 0}

	for sc.Scan() {
		number, _ := strconv.ParseInt(sc.Text(), 0, 64)
		if window[0] == 0  {
			window[0] = window[1]
			window[1] = window[2]
			window[2] = number
			fmt.Println(number, "(N/A - no previous sum)", window[0], window[1], window[2])
		} else {
			newsum = window[0] + window[1] + window[2]
			if newsum > oldsum {
				fmt.Println(newsum, "(increased)", window[0], window[1], window[2])
				increasecount++
			} else if newsum < oldsum {
				fmt.Println(newsum, "(decreased)", window[0], window[1], window[2])
			} else {
				fmt.Println(newsum, "(no change)", window[0], window[1], window[2])
			}
		}
		oldsum = newsum
		window[0] = window[1]
		window[1] = window[2]
		window[2] = number
	}
	fmt.Println("Final increase count:", increasecount)
}