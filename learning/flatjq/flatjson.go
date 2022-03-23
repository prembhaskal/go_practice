package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jeremywohl/flatten"
)

func main() {
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

	for k,v := range flatmap {
		fmt.Printf("%s  : %s\n", k, v)
	}
}
