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

	// get the users cmd choice
	var choice string
	_, err := fmt.Scanf("%s", &choice)
	check(err)

	// update current dir
	currentDir = currentDir + "/" + choice

	// display the chosen cmd doc
	f, err := os.Open(currentDir)
	check(err)

	b1 := make([]byte, 1024)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d, %s", n1, string(b1))
}
