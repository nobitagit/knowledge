## Short variable declaration (or the walrus operator)

```go

title := "Mr." // illegal

func main() {
  name := "John Green" // ok
}
```

- It can be used only inside functions. Every top level declaration need to **start with a keyword like var, const, func, type, and import**.
- Can't be used twice (no redeclare)
- Can be re-assigned though:

```go
name := "Tom"
name = "Nelson" // all fine, we're just reassigning a new value
name := "Gary" // error, we're redeclaring a new variable
```
