package main

import (
    "fmt"
    "net/http"
    "os"
    "./client"
    "./server/Domain"
)

func setupRoutes() {
       http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
            client.FileUploadForm(w, r)
       })

       http.HandleFunc("/api/v1/upload", func(w http.ResponseWriter, r *http.Request) {
            if server.FilesUpload(w, r) {
                fmt.Fprintf(w, "Successfully Uploaded File\n")
            }
       })
}

func main() {
    var PORT string
    if PORT = os.Getenv("PORT"); PORT == "" {
        fmt.Println("Port is not open!")
        os.Exit(1)
    }

    setupRoutes()

    http.ListenAndServe(":" + PORT, nil)
}
