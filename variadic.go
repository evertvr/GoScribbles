package main

import "fmt"





func main() {
  
  sl1 := []int{1,2,3,4,5}  

  sum(sl1...)

}




func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println("total = ", total)
}