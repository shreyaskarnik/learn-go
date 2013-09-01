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
    //learnTypes()                            // < y minutes, learn more!
}

// Functions can have parameters and (multiple!) return values.
func learnMultiple(x, y int) (sum, prod int) {
    return x + y, x * y // return two values
}
