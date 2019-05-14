package server

import (
    "fmt"
    "io"
    "os"
    "net/http"
    "encoding/base64"
    "mime/multipart"
)

// 10MB
const MAX_MEMORY = 10 * 1024 * 1024

func FileUpload(file *multipart.FileHeader) bool {
        fileHandler, err := file.Open()
        defer fileHandler.Close()
        if err != nil {
            fmt.Println(err)
            return false
        }

        newFileName := base64.StdEncoding.EncodeToString([]byte(file.Filename))
        out, err := os.Create("./storage/" + newFileName)
        defer out.Close()

        if err != nil {
             fmt.Println(err)
             return false
        }

        _, err = io.Copy(out, fileHandler)
        if err != nil {
            fmt.Println(err)
            return false
        }

        return true
}


func FilesUpload(w http.ResponseWriter, r *http.Request) bool {

   if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
       fmt.Println(err)
       http.Error(w, err.Error(), http.StatusForbidden)

   }

   formdata := r.MultipartForm
   files := formdata.File["file-input"]
   for i, _ := range files {
       FileUpload(files[i])
   }

   return true
}