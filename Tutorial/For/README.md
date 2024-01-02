# For

Unlike other lenguajes, `for` is the unique looping construct of the lenguaje.

### Examples
```go
// Classic initial/condition/after for loop
var i int

for i = 0; i < 4; i++ {
    fmt.Printf("%d ", i)
}

fmt.Println() // Jump line

// Same using the short variable declaration operator (':=')

for j := 0; j < 2; j++ {
    fmt.Println("How fun is Go")
}
```
### Output
```text
0 1 2 3
How fun is Go
How fun is go
```

However, it doesn't mean that we can't create `while` like loops.

```go
// While true
for {
    // do something ...
    break // Necessary to avoid get trapped
}

// Classic while loop

var j int = 4

for j > 0 {
    fmt.Printf("%d ", j)
}
```

