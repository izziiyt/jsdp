# Why DSTJ (Deep Sort on The Json)?

`dstj` is a simple tool to sort the all dictionaries and arrays in a JSON file. 
It is useful when you want to compare two JSON files, but the contents are not in the same order.
`dstj` is more powerful that it can sort arrays rather than `jq -S` or `dictknife` in some cases.

# Install

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
$ diff <(cat 0.json | go run . | jq) <(cat 1.json | go run . | jq)
6d5
<     false,
```
