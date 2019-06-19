// Hint : use ioutil.ReadFile

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
