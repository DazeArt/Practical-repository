package main
import "fmt"

func main() {
    var num int
    fmt.Scanf("%d", &num)
    var fivet int = num / 5000
    if fivet != 0{
        num -= fivet * 5000
    }
    var onet int = num / 1000
    if onet != 0 {
        num -= onet * 1000
    }
    var fiveh int = num / 500
    if fiveh != 0 {
        num -= fiveh * 500
    }
    var twoh int = num / 200
    if twoh != 0 {
        num -= twoh * 200
    }
    var oneh int = 0
    if num != 0 {
        oneh += num / 100
    }
    fmt.Println(fivet, onet, fiveh, twoh, oneh)
}