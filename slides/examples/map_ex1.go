package main

import "fmt"

func main() {
    states := make(map[string]string)
    states["AP"] = "Andhra Pradesh"
    states["KA"] = "Karnataka"
    states["MH"] = "Maharashtra"

    for code, name := range states {
	fmt.Printf("%s - %s\n", code, name)
    }
}