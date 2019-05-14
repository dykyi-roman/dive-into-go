package server

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/base64"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
      // Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
      r.ParseMultipartForm(10 << 20)
      // FormFile returns the first file for the given key `myFile`
      // it also returns the FileHeader so we can get the Filename, the Header and the size of the file
      file, handler, err := r.FormFile("file-input")

      if err != nil {
            fmt.Println("Error Retrieving the File")
            fmt.Println(err)
            return
      }
      defer file.Close()

      fmt.Printf("Uploaded File: %+v\n", handler.Filename)

      // Create a temporary file
      newFileName := base64.StdEncoding.EncodeToString([]byte(handler.Filename))
      tempFile, err := ioutil.TempFile("./storage", newFileName + "-")
      if err != nil {
            fmt.Println(err)
      }
      defer tempFile.Close()

      // read all of the contents of our uploaded file into a byte array
      fileBytes, err := ioutil.ReadAll(file)
      if err != nil {
           fmt.Println(err)
      }

      // write this byte array to our temporary file
      tempFile.Write(fileBytes)

      fmt.Fprintf(w, "Successfully Uploaded File\n")
}