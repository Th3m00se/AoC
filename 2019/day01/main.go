package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calcFuel(n int) int {
	x := (n/3 - 2)
	return x
}

func main() {
	//file, err := os.Open("testinput.txt")
	file, err := os.Open("inputs.txt")
	i := 0
	z := 0

	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, eachln := range text {
		// convert string module value
		x, err := strconv.Atoi(eachln)

		if err != nil {
			log.Fatal(err)
		}

		// calculate running total for fuel needed
		//i = calcFuel(x)
		//println("'i' is ", i)

		z = x

		for {
			z = calcFuel(z)
			//fmt.Println("'z' is ", z)
			if z <= 0 {
				break
			}
			i = i + z
		}
	}
	// print output value
	fmt.Println("total Fuel Required: ", i)
}
