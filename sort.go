package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
)

type SortedJSON struct {
	data map[string]interface{}
}

func NewSortedJSON(data map[string]interface{}) *SortedJSON {
	return &SortedJSON{data: data}
}

func (sj *SortedJSON) Sort() {
	sj.sortMap(sj.data)
}

// マップをソートする
func (sj *SortedJSON) sortMap(m map[string]interface{}) {
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			sj.sortMap(vv)
		case []interface{}:
			sj.sortSlice(vv)
		}
	}
}

// スライスをソートする
func (sj *SortedJSON) sortSlice(s []interface{}) {
	mpl := make([]map[string]interface{}, 0)
	sll := make([][]interface{}, 0)
	stl := make([]string, 0)
	bol := make([]bool, 0)
	fll := make([]float64, 0)
	nulll := make([]interface{}, 0)
	for _, v := range s {
		if v == nil {
			nulll = append(nulll, v)
			continue
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			sj.sortMap(vv)
			mpl = append(mpl, vv)
		case []interface{}:
			sj.sortSlice(vv)
			sll = append(sll, vv)
		case bool:
			bol = append(bol, vv)
		case string:
			stl = append(stl, vv)
		case float64:
			fll = append(fll, vv)
		default:
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	sort.StringSlice.Sort(stl)
	slices.SortFunc(fll, func(i, j float64) int {
		if i < j {
			return -1
		}
		return 1
	})
	slices.SortFunc(bol, func(i, j bool) int {
		x := 0
		if i {
			x++
		}
		if j {
			x--
		}
		return x
	})

	i := 0
	for _, v := range nulll {
		s[i] = v
		i++
	}
	for _, v := range fll {
		s[i] = v
		i++
	}
	for _, v := range bol {
		s[i] = v
		i++
	}
	for _, v := range stl {
		s[i] = v
		i++
	}
	for _, v := range sll {
		s[i] = v
		i++
	}
	for _, v := range mpl {
		s[i] = v
		i++
	}
}

// ソートされた JSON データをエンコードする
func (sj *SortedJSON) MarshalJSON() ([]byte, error) {
	return sj.marshalMap(sj.data)
}

// マップをソートしてエンコードする
func (sj *SortedJSON) marshalMap(m map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		v, err := sj.marshalValue(m[k])
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(&buf, "%q:%s", k, v)
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// 値をエンコードする
func (sj *SortedJSON) marshalValue(v interface{}) ([]byte, error) {
	switch vv := v.(type) {
	case map[string]interface{}:
		return sj.marshalMap(vv)
	case []interface{}:
		return sj.marshalSlice(vv)
	default:
		return json.Marshal(v)
	}
}

// スライスをエンコードする
func (sj *SortedJSON) marshalSlice(s []interface{}) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range s {
		if i > 0 {
			buf.WriteByte(',')
		}
		val, err := sj.marshalValue(v)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteByte(']')
	return buf.Bytes(), nil
}
