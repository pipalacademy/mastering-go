// Hint: use strings.ToUpper
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	fmt.Println(strings.ToUpper(arg))
}
