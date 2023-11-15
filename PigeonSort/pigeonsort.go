package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"sync"
)


func main(){
	start := time.Now()

	array := readCsv("in.csv")

	var wg sync.WaitGroup

	seperatedArraySize := len(array)/ 4

	for i:= 0;i <4; i++{
		start := i * seperatedArraySize
		end := (i+1) * seperatedArraySize

		if i == 3 {
			end = len(array)
		}

		wg.Add(1)
		go func(arr []int){
			
			defer wg.Done()
			pigeonSort(arr, 10000)
		}(array[start:end])
	}

	wg.Wait()

	pigeonSort(array, 10000)

	writeCsv(array)

	elapsed := time.Since(start)
	println()
	println()
	fmt.Println(elapsed)

}



func readCsv(name string) []int{
	var array []int
	file, err := os.Open(name)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil{
		panic(err)
	}

	for _, row := range data{
		for _, col := range row {
			stringValue, err := strconv.Atoi(col) 
			if err != nil{
				panic(err)
			}
			array = append(array, stringValue)
		}
	}

	return array
}


func writeCsv(array []int) {
	csvFile, err := os.Create("out.csv")

	if err != nil{
		panic(err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	
	for _, value := range array{
		stringVal := strconv.Itoa(value)
		err := csvwriter.Write([]string{stringVal})

		if err != nil{
			panic(err)
		}
	} 

	csvwriter.Flush()

	if err := csvFile.Sync(); err != nil{
		panic(err)
	}
	
}


func pigeonSort(array []int, length int) {

	min_val := array[0]
	max_val := array[0]

	for i := range array{
		if array[i]> max_val{
			max_val = array[i]
		}
		if array[i]< min_val{
			min_val = array[i]
		}
	}

	// fmt.Println(max_val)
	// fmt.Println(min_val)

	arraySize := max_val - min_val +1


	pigeonholes := make([]int, arraySize)

	
	for _, i := range array{
		pigeonholes[i - min_val]++
	}

	index := 0

	for i := 0; i < arraySize; i++ {
		for pigeonholes[i] > 0 {
			array[index] = i + min_val
			index++
			pigeonholes[i]--
		}
	}

}

