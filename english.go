package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var text string
	// use it as cmdline argument
	textArg := flag.String("text", "", "Text to search for")
	flag.Parse()
	// if cmdline arg was not passed ask
	if fmt.Sprintf("%s", *textArg) == "" {
		fmt.Print("Enter text: ")
		// get the sub string to search from the user
		fmt.Scanln(&text)
	} else {
		text = fmt.Sprintf("%s", *textArg)
	}
	// read the whole file at once
	b, err := ioutil.ReadFile("/home/xelpmoc/go/src/findword/Gujarati.txt")

	if err != nil {
		panic(err)
	}
	s := string(b)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, text))

	//counting number of repetations
	var count int = 0
	input := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range input {
		if v == text {
			if _, ok := m[v]; !ok {
				count = 1
				m[v] = count
			} else {
				count++
				m[v] = count
			}
		}
	}
	fmt.Println("count >>>> ", m[text])

	f, err := os.Open("/home/xelpmoc/go/src/findword/Gujarati.txt")
	if err != nil {
		return
	}
	defer f.Close()
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			fmt.Println(line)
			return
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}
