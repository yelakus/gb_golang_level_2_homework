package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func main() {

	recoverPanic()

	fmt.Println("Панику обработали и программа не завершилась аварийно")
}

func recoverPanic() {
	defer func() {
		value := recover()
		err, _ := value.(error)

		if err != nil {
			err = errors.Wrap(err, "Вызвано в recoverPanic()")
			log.Printf("recovered. %s", err)
		}
	}()

	doPanic()
}

func doPanic() {
	map1 := []int{5, 30, 50}
	fmt.Println(map1[4])
}
