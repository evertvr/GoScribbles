package main

import "fmt"

import "math"
import "time"


const s string = "constant string"

func main() {

/* declaring AND initting a var here */
    result := "bwrap"
    var e int32
    var f = true

    var a, b, c int = 1, 2, 3

    const n = 0.5
    const d = 3e20 / n


    fmt.Println(result)

    fmt.Println(f)

    fmt.Println(e)

    fmt.Println(a, "+",b,"+",c, "=", a+b+c)

    fmt.Println(n,d)

    fmt.Println(math.Sin(n))

    basicForLoop(4)

    t := time.Now()
    switch {
        case t.Hour() < 12:
        fmt.Println("it's before noon:",t.Hour())
        default:
        fmt.Println("it's after noon",t.Hour())
    }

    /* declare a array of 5 ints */
    var array1 [5]int

    /* declare AND initialize */
    array2 := [5]int{1,2,3,4,5}

    fmt.Println(array1, array2)


    fmt.Println("-------------------")

    sliceTest()

    fmt.Println("-------------------")
    mapTest()

    initMap := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", initMap)

}


func basicForLoop(maxIteration int){

    for i := 0; i <= maxIteration; i++ {
        fmt.Println(i)
    }

}

func sliceTest(){

    s := make([]string, 10)

    s[0] = "bwrap1"
    s[1] = "bwrap2"
    s[2] = "bwrap3"

    l := s[1:3]
    fmt.Println("sl1:", l)
    fmt.Println("l[0]:", l[0])
}

func mapTest(){

    m := make(map[string]int)


    m["leeftijd"] = 25
    m["lengte"]   = 180

    fmt.Println("map:", m)
    fmt.Println("lengte van map:",len(m))

    /* does the map contains the key "lengte" ? */

    _, lengtePresent := m["bwrap"]

    fmt.Println("does the map contains the key \"lengte\"?:",lengtePresent)
}