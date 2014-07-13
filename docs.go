package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	appName := os.Args[1]
	availableApps, _ := ioutil.ReadDir("./apps")
	selectedApp := ""

	for _, app := range availableApps {
		if app.Name() == appName {
			selectedApp = app.Name()
			fmt.Println("Selected", selectedApp)
		} else {
			fmt.Println("No docs available for", app)
		}
	}

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
		return
	}

}
