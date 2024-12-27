func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // Output: 0
    delete(m, "b")
    fmt.Println(m["b"]) // Output: 0
}