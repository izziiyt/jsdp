[![CI][ci-img]][ci]
[![Go Report Card][go-report-img]][go-report]
[![License: MIT][license-img]][license]

# Why DSTJ (Deep Sort on The Json)?

`dstj` is a simple tool for recursively sorting objects and arrays in JSON files.
It is useful in situations where you want to compare JSON files that have the same content but different order.
While `jq -S` and `dictknife` can sort objects, they cannot sort arrays.
In such cases, `dstj` is superior as it can also sort the contents of arrays.

# Install

(`go` is required)

```bash
go intsall github.com/izziiyt/dstj@v0.1.0
```

# Example

(`jq` is required)

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
$ diff <(cat 0.json | dstj | jq) <(cat 1.json | dstj | jq)
3d2
<     false,
```

# Sorting Order

Everything is sorted in ascending order.
The data types in the JSON are described according to [RFC-8259](https://datatracker.ietf.org/doc/html/rfc8259#section-3).

## object

Objects are sorted by the string order of their keys.

before
```json
{
  "c": false,
  "a": 1,
  "b": null
}
```

after
```json
{
  "a": 1,
  "b": null,
  "c": false
}
```

## array

Arrays are sorted first by the following order of data types:

- false
- true
- number
- string
- array
- object
- null

Subarrays divided by type are then sorted according to their respective types.

before
```json
{
  "a": [
    { "a":  1},
    false,
    "b",
    1,
    null,
    "a",
    [1, false],
    true,
    0.5
  ]
}
```

after
```json
{
  "a": [
    false,
    true,
    0.5,
    1,
    "a",
    "b",
    [false, 1],
    { "a":  1},
    null
  ]
}
```

[ci]: https://github.com/izziiyt/dstj/actions/workflows/ci.yaml
[ci-img]: https://github.com/izziiyt/dstj/actions/workflows/ci.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/izziiyt/dstj
[go-report-img]: https://goreportcard.com/badge/github.com/izziiyt/dstj
[license]: https://opensource.org/licenses/MIT
[license-img]: https://img.shields.io/badge/License-MIT-yellow.svg
