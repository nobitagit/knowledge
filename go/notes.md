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

While using the **var keyword** allows one to initialise the variable at a later stage after declaration, when using `:=` we are forced to initialise it with a value.
This means that the value/expression on the right will be used to infer the type, so we can leave it out.

It can be used to declare multiple variables all at once:

```go
person1, person2, person3 := 'Anthony', 'Jules', 'Hattie' // same type
person, age, isDev := 'Kyle', 23, false // different types
```

## Pointers

Read [this article](https://medium.com/rungo/pointers-in-go-a789eafccd53).
When you are declaring a variable, the Go compiler will allocate some memory for it in the RAM and depending on its type, it will allocate a specific size of memory to store the data.
That memory will have some **memory address** so that go can find that variable’s value when asked for it. These memory addresses are represented in hexadecimal values.

```sh
0x40e020
```

A pointer is a variable which points to the memory location of another variable.

```go
func main() {
	var x *int
	fmt.Printf("Type: %T\nValue: %v", x, x)
}
// Type: *int
// Value: <nil>
```

To access an address of a variable memory, go provides a simple operator & (ampersand) which if used in front of the variable name, and returns the memory address.
