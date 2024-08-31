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
	jsonData, err := io.ReadAll(os.Stdin)
	if err != nil {
		errExit(err)
	}

	var data map[string]any
	if err := json.Unmarshal(jsonData, &data); err != nil {
		errExit(err)
	}

	sortedJSON := NewSortedJSON(data)
	sortedJSON.Sort()

	sortedData, err := sortedJSON.MarshalJSON()
	if err != nil {
		errExit(err)
	}

	fmt.Println(string(sortedData))
}
