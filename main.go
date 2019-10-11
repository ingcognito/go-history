package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	histfile, err := os.Open("./hist1.txt")
	check(err)
	defer histfile.Close()
	fmt.Println("file successfully opened", histfile)

}

func readHistFile() {

	//look into NLP
	//Provide insights into most used binaries

	file, err := os.Open("./hist1.txt")
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
		for indexword, word := range words {
			found := false

			//Excluding histfile line number
			if indexword == 1 || word == " " || word == "" {
				found = true
				continue
			}

			for indexuwc, uwc := range uniqueWordCounters {

				if uwc.uniqueWord == word {
					//fmt.Println(uwc.uniqueWord, " = ", word)
					uniqueWordCounters[indexuwc].counter++
					//fmt.Println("what is uwc here: ", uwc)
					found = true
				}

			}

			if found == false {
				newUWC := uniqueWordCounter{uniqueWord: word, counter: 1}
				uniqueWordCounters = append(uniqueWordCounters, newUWC)
			}
		}
	}

	sort.Slice(uniqueWordCounters, func(i, j int) bool { return uniqueWordCounters[i].counter < uniqueWordCounters[j].counter })

	for _, uwc := range uniqueWordCounters {
		fmt.Println(uwc)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
