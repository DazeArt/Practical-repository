package main
import "fmt"

func main() {
    var x int
    fmt.Scanf("%d", &x)
    count := 0
    for x != 1 {
        if x % 2 == 0{
            x /= 2
        } else {
            x = 3 * x + 1
        }
        count += 1
    }
    fmt.Println(count)
}