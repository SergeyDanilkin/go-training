package main

import (
	"math/rand"
	"dice/output"
	"time"
)

const (
	cntDice = 2
	cntThrows = 100000
	min = 1
	max = 6
)

func main() {
	var numbers []int
	sums := make([]int, max * cntDice - 1)
	var procents []int
	var sum int
	var procent int

	rand.Seed(time.Now().UTC().UnixNano())

	for i:=min * cntDice; i <= max * cntDice; i++ {
		numbers = append(numbers, i)
		//sums = append(sums, 0)
	}

	for i:=0; i < cntThrows; i++ {
		sum = 0
		for k := 0; k < cntDice; k++{
			sum += random(min, max)
		}
		for j,num := range numbers {
			if num == sum {
				sums[j]++
				break
			}
		}
	}

	for i:=0; i < len(sums); i++ {
		procent = sums[i]*100/cntThrows
		procents = append(procents, procent)
	}
	output.PrintRow(numbers)
	output.PrintRow(sums)
	output.PrintRow(procents)
}

func random(min, max int) int {
    return rand.Intn(max - min + 1) + min
}