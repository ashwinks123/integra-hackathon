package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	// read file
	file, err := os.Open("./data/measurements.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tempMap := make(map[string][]string)
	byteArr := bufio.NewScanner(file)
	for byteArr.Scan() {
		line := strings.Split(byteArr.Text(), ";")
		if _, ok := tempMap[line[0]]; !ok {
			tempMap[line[0]] = make([]string, 0)
			tempMap[line[0]] = append(tempMap[line[0]], line[1])
		} else {
			tempMap[line[0]] = append(tempMap[line[0]], line[1])
		}
	}

	result := make([]string, len(tempMap))
	arrIdx := 0
	for city, temp := range tempMap {
		minVal, maxVal, sum, avg := 0.0, 0.0, 0.0, 0.0
		for _, val := range temp {
			parseFloat, _ := strconv.ParseFloat(val, 64)
			sum += parseFloat
			if parseFloat < minVal {
				minVal = parseFloat
			}
			if parseFloat > maxVal {
				maxVal = parseFloat
			}
		}
		avg = sum / float64(len(temp))
		result[arrIdx] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, minVal, avg, maxVal)
		arrIdx++
	}

	sort.Strings(result)
	fmt.Printf("{%s},", strings.Join(result, ","))
	fmt.Println(time.Since(startTime))

}
