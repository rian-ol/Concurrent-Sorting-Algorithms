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
	wg.Add(1)

	go cocktailMerge(array, &wg)

	wg.Wait()
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



func cocktailMerge(array []int, wg *sync.WaitGroup){
	defer wg.Done()
    size := len(array)
    swapped := true
    start := 0
    end := size - 1

    var mutex sync.Mutex

    for swapped == true {
        swapped = false

        var wgPass sync.WaitGroup
        wgPass.Add(2)

        go func() {
            defer wgPass.Done()
            for i := start; i < end; i++ {
                mutex.Lock()
                if array[i] > array[i+1] {
                    array[i], array[i+1] = array[i+1], array[i]
                    swapped = true
                }
                mutex.Unlock()
            }
        }()

        go func() {
            defer wgPass.Done()
            if !swapped {
                return
            }

            mutex.Lock()
            end = end - 1
            mutex.Unlock()

            for i := end - 1; i > start-1; i-- {
                mutex.Lock()
                if array[i] > array[i+1] {
                    array[i], array[i+1] = array[i+1], array[i]
                    swapped = true
                }
                mutex.Unlock()
            }
            start = start + 1
        }()

        wgPass.Wait()
    }
	
}