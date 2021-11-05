package main

import (
	"bufio"
	// "bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// USAGE - cat /var/tmp/test.out | go run learning/hextobin/hextobin.go > /var/tmp/data.gz
// remove newline from file ending if any using truncate -s -1 <<file>>
func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	hbr := hex.NewDecoder(br)

	_, err := io.Copy(bw, hbr)
	if err == io.EOF {
		// fmt.Fprintf(os.Stderr, "converted %d bytes to binary", n)
		return
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "error in converting to binary: %v", err)
		return
	}

	// fmt.Fprintf("converted %d bytes to binary", n)

	// manually doing it instead of io.copy
	// for {
	// 	rb := make([]byte, 1024)
	// 	wb := make([]byte, 1024)

	// 	n, err := br.Read(rb)
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Fprintf(os.Stderr, "error in read: %v\n", err)
	// 		break
	// 	}
	// 	// fmt.Fprintf(os.Stderr, "decode %d byes\n", n)

	// 	_, err = hex.Decode(wb, rb[0:n])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "error in hex decode: %v\n", err)
	// 		break
	// 	}

	// 	_, err = bw.Write(wb)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "write error : %v\n", err)
	// 		break
	// 	}
	// }

	// err := bw.Flush()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error flusing writer: %v", err)
	// }
}
