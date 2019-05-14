package main

import (
    "fmt"
    "net/http"
    "os"
    "./client"
    "./server/Domain"    
    "./server/Infrastructure" 
)

func setupRoutes() {
       http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {            
            form := Client.UploadForm{}    
            form.FileUploadForm(w, r)
       })

       http.HandleFunc("/api/v1/upload", func(w http.ResponseWriter, r *http.Request) {            
            var domain Domain.FileUploadInterface = Infrastructure.FileUploadLocalClient{}
            if domain.FilesUpload(w, r) {
                fmt.Fprintf(w, "Successfully Uploaded File(s)\n")
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
