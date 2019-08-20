// Every Go program is made of packages
package main // programs star running in 'main'

// Program is using import paths "fmt" and "math/rand"
import (
	"fmt"
	"math/rand"
)

func main() {
	// All exported names begin with capital letter
	// unexported names cannot be accessed outside the package
	fmt.Println("My favorite number is ", rand.Intn(10))
}
