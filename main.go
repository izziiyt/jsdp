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
	input, err := io.ReadAll(os.Stdin)
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
