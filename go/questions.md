- why colons are needed in maps for the last item (unless we avoid going on a new line)?

```go
// error
idMap := map[string]int {
    "Jane": 233
    }
// ok
idMap := map[string]int {
    "Jane": 233}
```

See also [here](https://stackoverflow.com/questions/9637483/syntax-error-unexpected-semicolon-or-newline-expecting)
