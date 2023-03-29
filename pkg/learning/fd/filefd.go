package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// how to use this file
//
//	$ tmpfilename=$(mktemp)
//	$ echo "this is some randome test data" > $tmpfilename
//	$ exec 3< $tmpfilename
//	$ rm $tmpfilename
//	$ go run filefd.go -fd 3
//	reading from file with fd: 3
//	read 31 bytes: 3
//	******** file contents **********
//	this is some randome test data
//
//	******** file contents end **********
//	$ cat <&3
//	$  <no output>
func main() {
	var filefd int
	flag.IntVar(&filefd, "fd", -1, "provide file fd")
	flag.Parse()

	if filefd == -1 {
		panic("fd is -1")
	}

	fmt.Printf("reading from file with fd: %d\n", filefd)

	tmpfile := os.NewFile(uintptr(filefd), "test file")
	defer func() {
		err := tmpfile.Close()
		if err != nil {
			fmt.Printf("error in file close: %v\n", err)
		}
	}()

	var sb strings.Builder
	n, err := io.Copy(&sb, tmpfile)
	if err != nil {
		fmt.Printf("error in file read: %v\n", err)
		return
	}

	fmt.Printf("read %d bytes\n", n)

	fmt.Println("******** file contents **********")
	fmt.Printf("%s", sb.String())
	fmt.Println("\n******** file contents end **********")
}
