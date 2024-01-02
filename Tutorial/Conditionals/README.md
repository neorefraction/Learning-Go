# Conditionals

Go provides us with `if`, `else if`, and `else` statements for branching and flow control, making them straightforward, like in many other programming languages.

### Example

```go
// If else statement
var name string

fmt.Print("Introduce your name: ")
fmt.Scan(&name)

if name == "Johnny" {
    fmt.Println("Welcome teacher!")
} else if name == "Carlos" {
    fmt.Println("Welcome dear Student!")
} else {
    fmt.Println("Unkown User. Access denied.")
}
```

### Output

```text
Introduce your name: Johnny
Welcome teacher!
```

Moreover, we can also use logical and relational operators to control the program flux.

### Example

```go
// Using logical and relational operators
var i, j int

// Asks for two numbers
fmt.Print("Introduce a number: ")
fmt.Scan(&i)

fmt.Print("Introduce another number: ")
fmt.Scan(&j)

// Evaluates the numbers
if i > 0 && j > 0 {
    fmt.Println("Both values are bigger than 0")
}

if i < 0 && j < 0 {
    fmt.Println("Both values are lower than 0")
}

if i == 0 && j == 0 {
    fmt.Println("Both values are lower than 0")
}

if i < 0 || j < 0 {
    fmt.Println("One of the values is lower than 0")
}
```

### Output

```text
Introduce a number: 3
introduce another number: -5
One of the values is lower than 0
```

On the other hand Go provides a switch structure too, to evaluate many diferent branches.

In summary, conditionals in Go are quite similar to those in other languages because, in the end, there isn't much more to add to this basic functionality. However, it's noteworthy to observe that parentheses `()` aren't needed, and that it isn't a ternary operator like in other high-level languages.