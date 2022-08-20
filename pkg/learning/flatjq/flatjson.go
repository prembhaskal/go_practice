package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jeremywohl/flatten"
)

func main() {
	var values bool
	flag.BoolVar(&values, "values", false, "to show values")
	flag.Parse()

	decoder := json.NewDecoder(os.Stdin)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	flatmap, err := flatten.Flatten(data, "", flatten.DotStyle)
	if err != nil {
		fmt.Printf("error in flatten: %v\n", err)
		return
	}

	for k, v := range flatmap {
		if values {
			fmt.Printf("%s  : %s\n", k, v)
		} else {
			fmt.Printf("%s\n", k)
		}
	}
}
