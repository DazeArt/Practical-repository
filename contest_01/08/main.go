package main
import "fmt"

func main() {
    var a int
    var b int
    var c int
    fmt.Scanf("%d %d %d", &a, &b, &c)
    if a > b && a > c{
        fmt.Println(a)
    }else if b > a && b > c{
        fmt.Println(b)
    }else {
        fmt.Println(c)
    }
}