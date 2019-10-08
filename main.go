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

type uniqueWordCounter struct {
	uniqueWord string
	counter    int
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
	uniqueWordCounters := make([]uniqueWordCounter, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		words := strings.Split(line, " ")

		for index, word := range words {
			found := false

			//Excluding histfile line number
			if index == 1 || word == " " {
				//For the histfile
				found = true
				continue
			}

			for _, uwc := range uniqueWordCounters {

				if uwc.uniqueWord == word {
					//fmt.Println(uwc.uniqueWord, " = ", word)
					uwc.counter++
					//fmt.Println("what is uwc here: ", uwc)
					found = true
				}

			}

			if found == false {
				newUWC := uniqueWordCounter{uniqueWord: word, counter: 1}
				uniqueWordCounters = append(uniqueWordCounters, newUWC)
			} //else {
			//fmt.Println("else is being hit")
			//}

		}
	}

	fmt.Println("Results: ", uniqueWordCounters)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
