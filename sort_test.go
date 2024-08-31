package main

import (
	"encoding/json"
	"testing"
)

const sampleJSON0 = `{
  "b": [[2, 0], "c", 1, {"a": 0}, 0.23, true, null, false],
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
		value    string
		name     string
		expected string
	}{
		{
			name:     "sampleJSON0",
			value:    sampleJSON0,
			expected: `{"a":1,"b":[false,true,0.23,1,"c",[0,2],{"a":0},null],"c":{"a":false,"b":2},"d":[1,{"a":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data map[string]any
			if err := json.Unmarshal([]byte(tt.value), &data); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			sortedData, err := NewSortedJSON(data).MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			res := string(sortedData)
			if string(sortedData) != tt.expected {
				t.Errorf("expected: %s, but: %s", tt.expected, res)
			}
		})
	}
}
