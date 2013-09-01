// Single line comment
/* Multi-
   line comment */

// A package clause starts every source file.
// Main is a special name declaring an executable rather than a library.
package main

// Import declaration declares library packages referenced in this file.
import (
    "fmt"      // A package in the Go standard library
    //"net/http" // Yes, a web server!
    //"strconv"  // String conversions
)

// A function definition.  Main is special.  It is the entry point for the
// executable program.  Love it or hate it, Go uses brace brackets.
func main() {
    // Println outputs a line to stdout.
    // Qualify it with the package name, fmt.
    fmt.Println("Hello world!")

    // Call another function within this package.
    beyondHello()
}

// Functions have parameters in parentheses.
// If there are no parameters, empty parens are still required.
func beyondHello() {
    var x int // Variable declaration.  Variables must be declared before use.
    x = 3     // Variable assignment.
    // "Short" declarations use := to infer the type, declare, and assign.
    y := 4
    sum, prod := learnMultiple(x, y)        // function returns two values
    fmt.Println("sum:", sum, "prod:", prod) // simple output
    learnTypes()                            // < y minutes, learn more!
}

// Functions can have parameters and (multiple!) return values.
func learnMultiple(x, y int) (sum, prod int) {
    return x + y, x * y // return two values
}
// Some built-in types and literals.
func learnTypes() {
    // Short declaration usually gives you what you want.
    s := "Learn Go!" // string type

    s2 := `A "raw" string literal
can include line breaks.` // same string type

    // non-ASCII literal.  Go source is UTF-8.
    g := 'Σ' // rune type, an alias for uint32, holds a UTF-8 code point

    f := 3.14195 // float64, an IEEE-754 64-bit floating point number
    c := 3 + 4i  // complex128, represented internally with two float64s

    // Var syntax with an initializers.
    var u uint = 7 // unsigned, but implementation dependent size as with int
    var pi float32 = 22. / 7

    // Conversion syntax with a short declaration.
    n := byte('\n') // byte is an alias for uint8

    // Arrays have size fixed at compile time.
    var a4 [4]int           // an array of 4 ints, initialized to all 0
    a3 := [...]int{3, 1, 5} // an array of 3 ints, initialized as shown

    // Slices have dynamic size.  Arrays and slices each have advantages
    // but use cases for slices are much more common.
    s3 := []int{4, 5, 9}    // compare to a3.  no ellipsis here
    s4 := make([]int, 4)    // allocates slice of 4 ints, initialized to all 0
    var d2 [][]float64      // declaration only, nothing allocated here
    bs := []byte("a slice") // type conversion syntax

    p, q := learnMemory() // declares p, q to be type pointer to int.
    fmt.Println(*p, *q)   // * follows a pointer.  This prints two ints.

    // Maps are a dynamically growable associative array type, like the
    // hash or dictionary types of some other languages.
    m := map[string]int{"three": 3, "four": 4}
    m["one"] = 1

    // Unused variables are an error in Go.
    // The underbar lets you "use" a variable but discard its value.
    _, _, _, _, _, _, _, _, _ = s2, g, f, u, pi, n, a3, s4, bs
    // Output of course counts as using a variable.
    fmt.Println(s, c, a4, s3, d2, m, bs)

    //learnFlowControl() // back in the flow
}
func learnMemory() (p, q *int) {
    // Named return values p and q have type pointer to int.
    p = new(int) // built-in function new allocates memory.
    // The allocated int is initialized to 0, p is no longer nil.
    s := make([]int, 20) // allocate 20 ints as a single block of memory
    s[3] = 7             // assign one of them
    r := -2              // declare another local variable
    return &s[3], &r     // & takes the address of an object.
}