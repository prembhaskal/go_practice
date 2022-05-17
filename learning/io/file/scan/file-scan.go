package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "", "specify file to parse")
	flag.Parse()

	parseUsingfscan(filename)
	fmt.Println("\n\n")
	parseUsingStringOps(filename)
}

func parseUsingStringOps(filename string) {
	parsefile, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("err in file open: %v", err))
	}

	defer parsefile.Close()

	bdata, err := ioutil.ReadAll(parsefile)
	if err != nil {
		panic(fmt.Sprintf("err in file read: %v", err))
	}

	sdata := string(bdata)
	lines := strings.Split(sdata, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue // ignore empty lines
		}

		words := strings.Split(line, " ")
		var first, second string
		first = words[0]
		if len(words) > 1 {
			second = words[1]
		}
		
		fmt.Printf("parsed output: first: %s, second: %s\n", first, second)
	}
}

func parseUsingfscan(filename string) {
	parsefile, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("err in file open: %v", err))
	}
	defer parsefile.Close()

	for i := 0; i < 8; i++ {
		var first, second string
		n, err := fmt.Fscanf(parsefile, "%s %s\n", &first, &second)
		fmt.Printf("fscan output, n: %d, err: %v\n", n, err)

		fmt.Printf("parsed output: first: %s, second: %s\n", first, second)
	}
}
