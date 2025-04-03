package playground

//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//	"strings"
//	"sync"
//	"time"
//)
//
//func main() {
//	startTime := time.Now()
//	// read file
//	file, err := os.Open("./data/measurements_1.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	//tempMap := make(map[string][]string)
//	byteArr := make([]byte, 1024)
//	fileReader := bufio.NewReader(file)
//	c := make(chan string)
//	wg := new(sync.WaitGroup)
//	//result := make([]string, 0)
//	for {
//		cnt, err := fileReader.Read(byteArr)
//		if err != nil && err == io.EOF {
//			break
//		}
//		//fmt.Println(string(byteArr[:cnt]))
//		wg.Add(1)
//		go processChunk(byteArr[:cnt-1], c, wg)
//		break
//		//result = append(result, string(byteArr[:cnt]))
//	}
//	go monitorWorker(wg, c)
//	//fmt.Println(result)
//	for msg := range c {
//		fmt.Println(msg)
//	}
//
//	//byteArr := bufio.NewScanner(file)
//	//for byteArr.Scan() {
//	//	line := strings.Split(byteArr.Text(), ";")
//	//	if _, ok := tempMap[line[0]]; !ok {
//	//		tempMap[line[0]] = make([]string, 0)
//	//		tempMap[line[0]] = append(tempMap[line[0]], line[1])
//	//	} else {
//	//		tempMap[line[0]] = append(tempMap[line[0]], line[1])
//	//	}
//	//}
//
//	//result := make([]string, len(tempMap))
//	//arrIdx := 0
//	//for city, temp := range tempMap {
//	//	min, max, sum, avg := 0.0, 0.0, 0.0, 0.0
//	//	for _, val := range temp {
//	//		parseFloat, _ := strconv.ParseFloat(val, 64)
//	//		sum += parseFloat
//	//		if parseFloat < min {
//	//			min = parseFloat
//	//		}
//	//		if parseFloat > max {
//	//			max = parseFloat
//	//		}
//	//	}
//	//	avg = sum / float64(len(temp))
//	//	result[arrIdx] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", city, min, avg, max)
//	//	arrIdx++
//	//}
//	//
//	//sort.Strings(result)
//	//fmt.Printf("{%s},", strings.Join(result, ","))
//	fmt.Println(time.Since(startTime))
//
//}
//
//func processChunk(chunk []byte, c chan string, wg *sync.WaitGroup) {
//	strChunk := string(chunk)
//	tempMap := make(map[string][]string)
//	for _, line := range strings.Split(strChunk, "\n") {
//		items := strings.Split(line, ";")
//		if _, ok := tempMap[items[0]]; !ok {
//			tempMap[items[0]] = make([]string, 0)
//			tempMap[items[0]] = append(tempMap[items[0]], items[1])
//		} else {
//			tempMap[items[0]] = append(tempMap[items[0]], items[1])
//		}
//	}
//	fmt.Println(tempMap)
//	c <- "Done chunk"
//	defer wg.Done()
//}
//
//func monitorWorker(wg *sync.WaitGroup, cs chan string) {
//	wg.Wait()
//	close(cs)
//}
