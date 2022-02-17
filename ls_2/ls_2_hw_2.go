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

		file, err := os.Create(FileName())

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}

		func() {
			defer file.Close()
		}()
	}

}

// FileName generates unique names with numbering for text files
func FileName() string {
	fileNum += 1
	return "empty_" + strconv.Itoa(fileNum) + ".txt"
}
