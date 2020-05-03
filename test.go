package main

import "fmt"
import "os"

// func main() {
// 	input := [] string {"flower","flow","flight"}

// 	fmt.Println(longerstCommonPrefix(input))
// }

const OUTER_COUNT = 4
const INNER_COUNT = 8
func main() {
    var s [][]string
    for i := 0; i < OUTER_COUNT; i++ {
        sl := make([]string,0,INNER_COUNT)
        for j := 0; j < INNER_COUNT; j++ {
            sl = append(sl,"a")
        }
        s = append(s,sl)
    }
    fmt.Println(s)
    os.Exit(1)
}

