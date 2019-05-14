package Domain

import (        
    "net/http"
)

// Interface
type FileUploadInterface interface{
    FilesUpload(w http.ResponseWriter, r *http.Request) bool
}