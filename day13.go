package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	println("Advent of code, day 13")

	filename := "puzzle.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	earliest, _ := strconv.Atoi(lines[0])
	busses := getBusses(lines[1])

	fmt.Println("earliest:", earliest)
	fmt.Println("busses:", busses)

	part1 := calcPart1(busses, earliest)
	fmt.Println("PART 1:", part1)

	part2 := calcPart2(busses)
	fmt.Println("PART 2:", part2)
}

func getBusses(line string) []int {
	var busses []int
	for _, busIdStr := range strings.Split(line, ",") {
		if busIdStr == "x" {
			busses = append(busses, 0)
		} else {
			busId, _ := strconv.Atoi(busIdStr)
			busses = append(busses, busId)
		}
	}
	return busses
}

func calcPart2(busses []int) int {
	largestIndex, largestBusId := getLargestBusId(busses)
	var step = largestBusId
	for timestamp := largestBusId - largestIndex; ; timestamp += step {
		bussesWithRemaindersZero := getBussesMatchingPart2(busses, timestamp)

		if len(bussesWithRemaindersZero) == len(busses) {
			return timestamp
		}

		step = calcNewStep(bussesWithRemaindersZero)
	}
}

func getBussesMatchingPart2(busses []int, timestamp int) []int {
	var bussesWithRemainders []int
	for idx, busId := range busses {
		if busId == 0 {
			bussesWithRemainders = append(bussesWithRemainders, 0)
		} else {
			rem := math.Mod(float64(timestamp+idx), float64(busId))
			if rem == 0 {
				bussesWithRemainders = append(bussesWithRemainders, busId)
			}
		}
	}
	return bussesWithRemainders
}

func getLargestBusId(busses []int) (int, int) {
	var largestId = 0
	var idx = 0
	for i, busId := range busses {
		if busId > largestId {
			largestId = busId
			idx = i
		}
	}
	return idx, largestId
}

func calcNewStep(bussesWithRemaindersZero []int) int {
	var step = 1
	for _, busId := range bussesWithRemaindersZero {
		if busId != 0 {
			step *= busId
		}
	}
	return step
}

func calcPart1(busses []int, earliest int) int {
	minutesPast := getBussesWithMinutesPast(busses, earliest)
	earliestBus := getEarliestBus(minutesPast)
	part1 := earliestBus[0] * earliestBus[1]
	return part1
}

func getBussesWithMinutesPast(busses []int, earliest int) [][]int {
	var minutesPast [][]int
	for _, it := range busses {
		if it != 0 {
			minutePast := int(math.Mod(
				float64((earliest/it+1)*it-earliest),
				float64(it)))
			minutesPast = append(minutesPast, []int{it, minutePast})
		}
	}
	return minutesPast
}

func getEarliestBus(minutesPast [][]int) []int {
	var min = []int{0, math.MaxInt64}
	for _, minutePast := range minutesPast {
		if minutePast[1] < min[1] {
			min = minutePast
		}
	}
	return min
}
