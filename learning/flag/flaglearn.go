package main

import (
	"flag"
	"fmt"
)

// below is the result
//  $ go run learning/flag/flaglearn.go  --leader-elect
// value of flag enable Leader: true
//  $ go run learning/flag/flaglearn.go  --leader-elect false
// value of flag enable Leader: true
//  $ go run learning/flag/flaglearn.go  --leader-elect=false
// value of flag enable Leader: false

func main() {
	var enableLeader bool
	flag.BoolVar(&enableLeader, "leader-elect", false, "to enable leader elect")
	flag.Parse()

	fmt.Printf("value of flag enable Leader: %v\n", enableLeader)
}
