[![CI][ci-img]][ci]
[![Go Report Card][go-report-img]][go-report]
[![License: MIT][license-img]][license]

# Why DSTJ (Deep Sort on The Json)?

`dstj` is a simple tool to sort the all objects and arrays in a JSON file. 
It is useful when you want to compare two JSON files, but the contents are not in the same order.
`dstj` is more powerful that it can sort arrays rather than `jq -S` or `dictknife` in some cases.

# Install

(`go` is required to install `dstj`)

```bash
go intsall github.com/izziiyt/dstj@v0.1.0
```

# Example

```bash
$ cat 0.json
{
  "b": 2,
  "a": [
    false,
    { 
      "a": 1 
    },
    1,
    null,
    2.1,
    "value",
    [3, 2, 1]
  ]
}
$ cat 1.json
{
  "a": [
    2.1,
    1,
    { 
      "a": 1 
    },
    null,
    [2, 3, 1],
    "value"
  ],
  "b": 2
}
$ diff <(cat 0.json | dstj . | jq) <(cat 1.json | dstj . | jq)
3d2
<     false,
```

# Sorting Order

以下は昇順 `--ascending=true` でソートした場合の解説をします。json 内のデータ型は [RFC-8259](https://datatracker.ietf.org/doc/html/rfc8259#section-3) の表記に従います

## object

object はキーの文字列順序でソートされます。

## array

array はまず以下の順序でデータの型に応じてソートされます。

- false
- true
- number
- string
- array
- object
- null

型で分けられた部分配列はそれぞれの型に応じたソートがされます。

[ci]: https://github.com/izziiyt/dstj/actions/workflows/ci.yaml
[ci-img]: https://github.com/izziiyt/dstj/actions/workflows/ci.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/izziiyt/dstj
[go-report-img]: https://goreportcard.com/badge/github.com/izziiyt/dstj
[license]: https://opensource.org/licenses/MIT
[license-img]: https://img.shields.io/badge/License-MIT-yellow.svg