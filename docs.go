package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	apps, _ := ioutil.ReadDir("./apps")

	var count = 1
	for _, app := range apps {
		fmt.Println(count, app.Name())
		count++
	}
}
