package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func errExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	var input []byte
	var err error

	if len(os.Args) > 1 {
		input, err = os.ReadFile(os.Args[0])
	} else {
		// if argument is not passed, read data from stdin
		input, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		errExit(err)
	}

	var data map[string]any
	if err := json.Unmarshal(input, &data); err != nil {
		errExit(err)
	}

	res, err := NewSortedJSON(data).MarshalJSON()
	if err != nil {
		errExit(err)
	}

	fmt.Println(string(res))
}
