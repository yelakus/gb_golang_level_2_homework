package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	fileNum int
)

func main() {
	for i := 0; i < 1000000; i++ {

		file, err := os.Create(fileName())

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}

		func() {
			defer file.Close()
		}()
	}

}

func fileName() string {
	fileNum += 1
	return "empty_" + strconv.Itoa(fileNum) + ".txt"
}
