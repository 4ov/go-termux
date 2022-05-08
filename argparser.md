# Refrence for Arguments parser [major.go](./major.go#L10)

## Supported data types

### `string`
```go
Name string `arg:"--name"` // [--name string]
```


### `bool`
```go
MilkFree bool `arg:"--use-milk"` // [--use-milk]`
```

### `int`
```go
Rounds int `arg:"-r"` // [-r 10]
```

### `[]Type` (aka slice)
```go
ages []int `arg:"--ages" split:","` // [--ages 1,2,3]
//or
ages []string `arg:"--names" split:","`
```