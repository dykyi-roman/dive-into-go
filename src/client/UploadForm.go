package Client

import (
    "net/http"
    "html/template"
)

const FileUploadTemplatePath string = "client/Templates/index.html"

type ViewData struct {
    Title string
}

type UploadForm struct{}

func (f UploadForm) FileUploadForm(w http.ResponseWriter, r *http.Request) {
     data := ViewData {
        Title: "vasvMulti upload file form",
     }

     tmpl, _ := template.ParseFiles(FileUploadTemplatePath)
     tmpl.Execute(w, data)
}