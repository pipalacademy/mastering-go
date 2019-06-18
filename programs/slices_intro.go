package main

import "fmt"

func main() {
	var langs = []string{"go", "python", "c", "java", "haskell"} // slice

	fmt.Printf("%T %v %v %v\n", langs, langs, cap(langs), len(langs))

	favLangs := langs[0:3]
	cpy_langs := langs
	fmt.Printf("%T %v %v %v\n", favLangs, favLangs, cap(favLangs), len(favLangs))

	favLangs[2] = "rust"
	fmt.Println(favLangs)
	fmt.Println(langs)
	fmt.Println(cpy_langs)
}
