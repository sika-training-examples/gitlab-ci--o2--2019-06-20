package main
import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Awesom Hello World from O2!\n")
}

func main() {
    http.HandleFunc("/", index)
    fmt.Println("Server startded.")
    http.ListenAndServe(":80", nil)
}