package main
import (
 "fmt"
 "net/http"
 "os"
)

func main() {
 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
   fmt.Println("Port is not open!")
   os.Exit(1)
 }

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello from path: %s\n", r.URL.Path)
 })

 http.ListenAndServe(":" + PORT, nil)
}
