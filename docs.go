package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	const dir = "./apps"
	currentDir := ""

	appName := os.Args[1]
	availableApps, _ := ioutil.ReadDir(dir)

	// select an available app
	for _, app := range availableApps {
		if app.Name() == appName {
			currentDir = dir + "/" + appName
			// fmt.Println(currentDir)
			break
		}
	}

	// app not found
	if currentDir == "" {
		fmt.Println(appName, "not found ...")
		return
	}

	// load cmd choices
	availableCmd, _ := ioutil.ReadDir(currentDir)
	for _, cmd := range availableCmd {
		fmt.Println(">", cmd.Name())
	}

	fmt.Print("select a cmd: ")

	// get the users cmd choice
	var choice string
	_, err := fmt.Scanf("%s", &choice)
	check(err)

	// update current dir
	currentDir = currentDir + "/" + choice

	// display the chosen cmd doc
	b, err := ioutil.ReadFile(currentDir)
	check(err)
	fmt.Printf("> %s", string(b))
}
