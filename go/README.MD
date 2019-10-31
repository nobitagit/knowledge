# Go

- compiled
- automatically memory management (garbage collection)
- OOP
- does not have the concept of class

## Structs
Instead of `class`es Go has `struct`s.
The are similar, but with some key limitation:

- no inheritance
- no constructors
- no generics

## Concurrency
Moore's law says that the number of transistors doubles every 18 months.
This meant that computers have been getting consistently faster every year. This used to be tru, but reality is now different. The development has hit limitations such as power/temperature constraints.
Heat is one major road block, since transistors risk melting as temperatures increase.

A solution to this has been increasing the number of cores. Multiple cores means that different tasks can be performed on the same machine at the same time (parallelism).
Parallelism is inherently hard. Some question arise:

- when do tasks start/stop?
- what if one task need data from another parallel task?
- how do they handle conflicts in memory. For ex. how do we make sure that one process writing to variable A does not conflict to another process writing to variable A?

Concurrency is key for large systems, and Go has primitives for it:

- **goroutines**: concurrent tasks
- **channels**: used to communicate between tasks
- **select**: to handle task synchronisation

## Workspaces
All go files are stored in workspaces, with a predefined hierarchy of directories.
This convention favors sharing code, which is one of the most important points in Go.

Workspaces are divided in 3 subdirectories:
- **src**: for source code 
- **pkg**: for libs
- **bin**: for executables

The unique workspace directory is defined by the `GOPATH` environment variable.
This is set automatically on installation.
The default one should be:

```sh
~/go
```

All go tools assume the Go code is inside `GOPATH`.

## Packages
The first line of each file gives the name to the package.

```go
package dateutils
```
```go
package currencyutils
```

These will be imported like so:

```go
import (
  "dateutils"
  "currencyutils"
)
```

Execution starts on package `main`, which is required.
Upon compilation this file will be the one that will be turned into the executable.
The main package needs a function called `main`.

```go
package main
import "fmt"

func main() {
  fmt.Printf("hello")
}
```

## [Go Tool](https://www.coursera.org/learn/golang-getting-started/lecture/09AcC/m1-2-3-go-tool)
Go tool is downloaded with Go itself and includes a series of commands:

- **`go build`**: compiles the program.
- **`go fmt`**: formats code
- **`go get`**: package manager
- **`go run`**: compiles go files and runs the executable
- **`go test`**: runs tests

## [Variables](https://www.coursera.org/learn/golang-getting-started/lecture/dQajs/m1-3-1-variables)

- case sensitive
- keywords are forbidden
- need a type declaration

```go
var x int
```

The type defines which values a variable accepts and which operations this variable supports.
For ex. if a variable is an integer it will support addition, division etc.

Variable types are initialised on declaration. Type can also be inferred, although it's always preferable to be explicit

```go
var age int = 22
var height = 177 // int is inferred
```

Values can be assigned to variables after declaration:

```go
var age int
age = 22
```

Unitialised variables are automatically assigned the "zero" value of that type:

```go
var age int // 0
var name string // ""
```

Variables can also be declared and initialised via **short variable declaration**:

```go
age := 22
```
Type is inferred from the right hand side. Short variable declarations can only be used **inside** functions.

## Data types

### [Pointers](https://www.coursera.org/learn/golang-getting-started/lecture/uYnqu/m2-1-1-pointers)
A pointer is an address to some data in memory. 2 operators are associated with pointers:

 1. `&`: returns the address of a variable/function
 2. `*`: returns data at an address (dereferencing) 

So they are the opposite of one another.

```go
var x int = 1
var y int
var p *int // p is a pointer to an integer (any integer)

p = &x // take the address in memory where the value 1 is stored and assign this reference to p
y = *p // take the value that p references to and assign it to y
```
### New
New is a function, used to create a variable and return a pointer to that variable.
The variable is initialised to 0.

```go
num := new(int)
*num = 1 // set the value of that integer
```
