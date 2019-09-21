// Each folder that contains `.go` source code files needs to be in a "package".
// A package is made up of one or more Go source code files. All files within the
// same package can access the functions and global variables (even "unexported" ones)
// within the package, even if that function or global variable is in a different
// `.go` file.
package main

// The import keywords allows you to include other packages, and are namespaced to
// the package name. Third party packages are imported using the full import path -
// e.g. "github.com/<user>/<repo>". Some packages are given vanity URLs, such as
// "k8s.io/apimachinery".
import "fmt"

func main() {
	// Go is strongly and statically typed, but has type inference via the := operator.
	// As a result the compiler will infer the type based on what is on the right-hand
	// side of the := operator. In this case it will infer it is a string.
	friend := "DDD East Anglia"

	// We are using the `fmt` package here to access the `Printf` function.
	// `fmt` is part of the standard library, and as the name suggests provide
	// functions for formatting and printing strings.
	fmt.Printf("Hello, %s", friend)
}
