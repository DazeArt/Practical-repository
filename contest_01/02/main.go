package main
import "fmt"

func main() {
    var ox_man float32 = 365.0/2.0
    var topol int = int(ox_man) / 32 + 1
    var dub int = int(ox_man) / 20 + 1
    fmt.Printf("%.1f %d %d", ox_man, topol, dub)
}