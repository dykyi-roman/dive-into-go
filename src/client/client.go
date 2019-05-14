package client

import (
    "net/http"
    "html/template"
)

const FileUploadTemplatePath string = "client/templates/index.html"

type ViewData struct{
    Title string
}

func FileUploadForm(w http.ResponseWriter, r *http.Request) {
     data := ViewData {
        Title: "Multi upload file form",
     }

     tmpl, _ := template.ParseFiles(FileUploadTemplatePath)
     tmpl.Execute(w, data)
}