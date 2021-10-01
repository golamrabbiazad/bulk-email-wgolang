package main

import (
	statenames "bulk-email/src/stateNames"
	"fmt"
)

func main() {
	allState := statenames.StateNames()

	for idx, val := range allState {
		fmt.Println(idx, val)
	}
}
