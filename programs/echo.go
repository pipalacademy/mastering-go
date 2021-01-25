package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
}
