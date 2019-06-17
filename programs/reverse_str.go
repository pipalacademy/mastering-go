package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func strjoin(sep string, strs []string) string {
	var r string
	for i, s := range strs {
		if i != 0 {
			r += sep
		}
		r += s
	}
	return r
}

func main() {
	s := "string"
	fmt.Println(s, reverse(s))

	s = "œ∑®†¥"
	fmt.Println(s, reverse(s))

	fmt.Println(strjoin(",", []string{"hello", "world"}))
}
