# AnyMap

## Usage

```go
am := anymap.New(...)

// get value
am.ObjectOf("user").ObjectOf("name").Value().(string)
// or get string with default value
am.ObjectOf("user").ObjectOf("name").String("")
```