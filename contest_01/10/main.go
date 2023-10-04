package main
import "fmt"

func main() {
    var a string
    var b string
    var c string
    
    fmt.Scanf("%s %s %s", &a, &b, &c)
    if a == "Да" && b == "Да" && c == "Да"{
        fmt.Println("Утка")
    } else if a == "Да" && b == "Да" && c == "Нет"{
        fmt.Println("Пингвин")
    } else if a == "Да" && b == "Нет" && c == "Да"{
        fmt.Println("Плезиозавры")
    } else if a == "Да" && b == "Нет" && c == "Нет"{
        fmt.Println("Дельфин")
    } else if a == "Нет" && b == "Да" && c == "Да"{
        fmt.Println("Страус")
    } else if a == "Нет" && b == "Да" && c == "Нет"{
        fmt.Println("Курица")
    } else if a == "Нет" && b == "Нет" && c == "Да"{
        fmt.Println("Жираф")
    } else if a == "Нет" && b == "Нет" && c == "Нет"{
        fmt.Println("Кот")
    }
    
}