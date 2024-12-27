func main() {
    var m = make(map[string]int)
    m["a"] = 1
    value, ok := m["b"]
    if ok {
        fmt.Println("Value of b:", value)
    } else {
        fmt.Println("Key b does not exist in the map")
    }
    delete(m, "b")
    value, ok = m["b"]
    if ok {
        fmt.Println("Value of b:", value) 
    } else {
        fmt.Println("Key b does not exist in the map")
    }
} 