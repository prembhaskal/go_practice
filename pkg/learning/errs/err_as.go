package main

import (
	"errors"
	"fmt"
)

func ErrorsAsPractice() {
	e2 := &error2{}
	wrape1 := newerror1(e2)

	fmt.Printf("error happened: %s\n", wrape1)

	var orige2 *error2
	iswrapped := errors.As(wrape1, &orige2)
	if iswrapped {
		fmt.Println("********** yes wrapped error *************")
		fmt.Printf("orig error: %s\n", orige2)
	}
}

type error1 struct {
	orig error
}

func newerror1(orig error) error {
	return fmt.Errorf("wrapped erorr: %w", orig)
	// return &error1{orig: orig}
}

func (e1 error1) Error() string {
	return "this is error1"
}

type error2 struct {

}

func (e2 error2) Error() string  {
	return "this is error2"
}