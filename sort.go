package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

type SortedJSON struct {
	data map[string]any
}

func NewSortedJSON(data map[string]any) *SortedJSON {
	sj := &SortedJSON{data: data}
	sj.sortMap(sj.data)
	return sj
}

func (sj *SortedJSON) sortMap(m map[string]any) {
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]any:
			sj.sortMap(vv)
		case []any:
			sj.sortSlice(vv)
		}
	}
}

// スライスをソートする
// スライスの要素がマップやスライスの場合は再帰的にソートする
// スライスの要素が数値、真偽値、文字列の場合はそれぞれのスライスに分けてソートする
// nil はスライスの最後に移動する
// スライスの要素が異なる型の場合は順序が
func (sj *SortedJSON) sortSlice(s []any) {
	blsl := make([]bool, 0)
	flsl := make([]float64, 0)
	stsl := make([]string, 0)
	slsl := make([][]any, 0)
	mpsl := make([]map[string]any, 0)
	nlsl := make([]any, 0)

	// スライスの要素を型ごとに分ける
	for _, v := range s {
		switch vv := v.(type) {
		case bool:
			blsl = append(blsl, vv)
		case float64:
			flsl = append(flsl, vv)
		case string:
			stsl = append(stsl, vv)
		case []any:
			sj.sortSlice(vv)
			slsl = append(slsl, vv)
		case map[string]any:
			sj.sortMap(vv)
			mpsl = append(mpsl, vv)
		case nil:
			nlsl = append(nlsl, vv)
		default:
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}

	// スライスを型ごとにソートする
	sort.StringSlice.Sort(stsl)
	sort.Float64s(flsl)
	sort.Slice(blsl, func(i, j int) bool { return !blsl[i] && blsl[j] })

	// 元のスライスにソートされた要素を再配置
	i := 0
	for _, v := range blsl {
		s[i] = v
		i++
	}
	for _, v := range flsl {
		s[i] = v
		i++
	}
	for _, v := range stsl {
		s[i] = v
		i++
	}
	for _, v := range slsl {
		s[i] = v
		i++
	}
	for _, v := range mpsl {
		s[i] = v
		i++
	}
	for _, v := range nlsl {
		s[i] = v
		i++
	}
}

func (sj *SortedJSON) MarshalJSON() ([]byte, error) {
	return sj.marshalMap(sj.data)
}

func (sj *SortedJSON) marshalMap(m map[string]any) ([]byte, error) {
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

func (sj *SortedJSON) marshalValue(v any) ([]byte, error) {
	switch vv := v.(type) {
	case map[string]any:
		return sj.marshalMap(vv)
	case []any:
		return sj.marshalSlice(vv)
	default:
		return json.Marshal(v)
	}
}

func (sj *SortedJSON) marshalSlice(s []any) ([]byte, error) {
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
