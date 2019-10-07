package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type histFile struct {
	line     int
	commands string
}

func main() {
	fmt.Printf("hello world")
	readHistFile()
}

func getHistFile() {
	//Get Histfile by using environment variable
	//Allow user to enter their own histfile they would like analyzed

	histfile, err := os.Open("./hist.txt")
	check(err)
	defer histfile.Close()
	fmt.Println("file successfully opened", histfile)

}

func readHistFile() {

	file, err := os.Open("./hist.txt")
	check(err)
	defer file.Close()

	var lines []string
	var git int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//this allows to read the lines, line by line, this is a ghetto analyzer but can pump out stats
	fmt.Println("read lines:")
	for _, line := range lines {
		if strings.Contains(line, "git") {
			git++
		}
	}
	fmt.Println("there are this many lines:", len(lines))
	fmt.Println("first line: ", lines[0])
	fmt.Println("you used git this many times: ", git)
	// most used binary
	// most used languages
	// date of of first terminal usage
	//
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
