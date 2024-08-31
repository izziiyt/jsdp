package main

import (
	"encoding/json"
	"testing"
)

const sampleJSON0 = `{
  "b": [[2, 0], "c", 1, {"a": 0}, 0.23, true, null],
  "a": 1,
  "d": [
    {
      "a": 1
    },
    1
  ],
  "c": {
    "b": 2,
    "a": false
  }
}
`

func TestSort(t *testing.T) {
	tests := []struct {
		value string
		name  string
		exp   string
	}{
		{
			name:  "sampleJSON0",
			value: sampleJSON0,
			exp:   `{"a":1,"b":[null,0.23,1,true,"c",[0,2],{"a":0}],"c":{"a":false,"b":2},"d":[1,{"a":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var data map[string]any
			if err := json.Unmarshal([]byte(tt.value), &data); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			sortedJSON := NewSortedJSON(data)
			sortedJSON.Sort()

			sortedData, err := sortedJSON.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			res := string(sortedData)
			if string(sortedData) != tt.exp {
				t.Errorf("want %s but %s", tt.exp, res)
			}
		})
	}
}
