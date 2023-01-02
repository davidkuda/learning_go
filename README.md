# Learning Go

## Book Notes "The Go Programming Language"

### Chapter 2: Program Structure

- page 28: _"convention and style in Go programs lean toward short names, especially for local variables with small scopes. Generally, the large the scope of a name, the longer and more meaningful it should be."_
- page 28: _"There are four major kind of declarations: var, const, type and func."_
- page 30: `var name type = expression` -- either type or expression may be omitted
- page 30: initial values:
  - pointers: nil
  - int: 0
  - booleans: false
  - strings: ""
  - interfaces: nil
  - reference types: nil (slice, pointer, map, channel, function)
  - aggregate types: zero value of all of its elements or fields (arrays or structs)
- page 30: `value := 108` -- Short Variale Declaration (possible within function scope)
- page 31: _"Keep in mind that := is a declaration, whereas = is an assignment."_
- page 31:
  ```go
  // first err is a declaration
  in, err := os.Open(inFile)
  // second err is an assigment (whereas out is a declaration)
  out, err := os.Create(outFile)
  // at least one of the list must be a new declaration, otherwise: compile error
  ```
- page 32:
  - A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. [...] With a pointer, we can read or update the value of a variable indirectly, without using or even knowing the name of the variable, if indeed it has a name.
  - If a variable is declared var x int, the expression **&x (“address of x”)** yields a pointer to an integer variable, that is, a value of type **\*int, which is pronounced “pointer to int.”** If this value is called p, we say “p points to x,” or equivalently “p contains the address of x.” The variable to which p points is written *p. The expression *p yields the value of that variable, an int, but since \*p denotes a variable, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable.
- page 32: Same, but spread / sparse:
  - **&x (“address of x”)** yields a pointer to an integer variable
  - **\*int, which is pronounced “pointer to int.”**
  - `var p *int = &num` -- Declare the variable p as a pointer to an integer that holds the address of num
  - “p points to x” or equivalently “p contains the address of x”
  - The expression \*p yields the value of that variable

### Chapter 4: Composite Types

#### Maps

```go
ages := map[string]int{
        "David": 30,
        "Elon":  30,
    }

ages := make(map[string]int)
ages["Richard"] = 30

ages["John"] // == 0 (default value of int), because not in map

age, ok := ages["Lennon"] 
if !ok { /* handle */ }

// combine both statements in one line:
age, ok := ages["Paul"]; !ok { /* handle */ }

// most operations on a nil map will work:
var nilMap map[string]int
nilMap["McCartney"] // = 0
// but storing to a nil map causes a panic:
nilMap["McCartney"] = 30 // panic: Assignment to entry in nil map
// You must allocate the map before you can store into it.
```

```go
```

#### Structs

Sample struct:

```go
type person struct {
    name string
    age  int
}
```

Three ways of specifying struct values:

```go
var david person
david.name = "David"
david.age = 30

// Struct Literals 1:

// All fields must be assigned a value, right order must be kept
david := person{"David", 30}

// Struct Literals 2 ("named"):

// order is not important.
// Completeness neither: Set to corresponding zero value, if not defined.
david := person{name: "David"}
```

Struct Embedding (page 105):

```go
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle // Anonymous Field: no name; makes access to values easier.
    Spokes int
}

w := Wheel{
    Circle: Circle {
        Point: Point {X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 21,
}
```

### Pointers

```go
var val int = 108

// a pointers stores the address of a variable
var p *int = &val // & retrieves the address of a variable

// p points to val
// p has the type *int, which means ...

// Dereferencing a  pointer
*p++ // retrieve value of p
```
