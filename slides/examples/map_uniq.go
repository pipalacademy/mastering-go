// Program uniq prints all the unique words given the command line arguments
// USAGE:
//    $ go run unique.go one two three two one
//    one
//    two
//    three
package main
import "fmt"
import "os"

func main() {
    // take the words from command-line arguments
    words := os.Args[1:]

    // make a map to remember the already seen words
    seen := make(map[string]bool)

    for _, w := range words {
        if !seen[w] {       // the word is not seen yet?
            fmt.Println(w)  // print it
            seen[w] = true  // and mark it as seen
        }
    }
}