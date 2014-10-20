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

	fmt.Println(" ")

	// const dir = "./apps" - I need a way for `go install` to take the relevant directories with it
	const dir = "/home/richard/dev/go/src/github.com/richarddowner/docs/apps"
	currentDir := ""

	appName := os.Args[1]
	availableApps, _ := ioutil.ReadDir(dir)

	// match app selection
	for _, app := range availableApps {
		if app.Name() == appName {
			currentDir = dir + "/" + appName
			break
		}
	}

	// app not found
	if currentDir == "" {
		fmt.Println(appName, "not found ...")
		return
	}

	// load cmd choices
	var count = 1
	commands := make(map[int]string)
	availableCmd, _ := ioutil.ReadDir(currentDir)
	for _, cmd := range availableCmd {
		fmt.Printf("(%d) %s\n", count, cmd.Name())
		commands[count] = cmd.Name()
		count++
	}

	fmt.Println("")
	fmt.Print(">> select: ")

	// get the users cmd choice
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	check(err)

	// map choice to app
	chosenApp := string(commands[choice])

	// update current dir
	currentDir = currentDir + "/" + chosenApp

	// display the chosen cmd doc
	b, err := ioutil.ReadFile(currentDir)
	check(err)
	fmt.Printf("%s", string(b))
}
