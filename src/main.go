package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/dykyi-roman/dive-into-go/client"    
    "github.com/dykyi-roman/dive-into-go/server/Domain"
    "github.com/dykyi-roman/dive-into-go/server/Infrastructure"    
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
