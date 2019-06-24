package main

import "fmt"

// calculate total and put it in channel

func sum(s []int, c chan int) {
    total := 0
    for _, v := range s {
        total += v
    }
    c <- total
}

func concurrentSum(nums []int, numChannels int) int {
    partial, total := 0, 0
    start, end := 0, 0
    c := make(chan int)
    chunkSize := len(nums) / numChannels

    for i := 0; i < numChannels; i++ {
        end = start + chunkSize
        go sum(nums[start:end], c)
        start += chunkSize
    }
    for i := 0; i < numChannels; i++ {
        partial = <-c
        total += partial
    }
    /* if we combine the above 2 for loops, it'd still work, but not totally concurrently.
       Think why is it so?
    */

    return total
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // concurrentSum handles all the go routine handling
    fmt.Println(concurrentSum(nums, 2))
}
