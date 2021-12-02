package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var oldnumber int64 = -1
	increasecount := 0

	for sc.Scan() {
		number, _ := strconv.ParseInt(sc.Text(), 0, 32)
		if oldnumber < 0 {
			fmt.Println(number, " (N/A - no previous measurement)")
		} else if number > oldnumber {
			fmt.Println(number, " (increased)")
			increasecount++
		} else if number < oldnumber {
			fmt.Println(number, " (decreased)")
		} else {
			fmt.Println(number, " (same)")
		}
		oldnumber = number
	}
	fmt.Println("Final increase count:", increasecount)
}