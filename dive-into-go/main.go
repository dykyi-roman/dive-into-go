package main

import (
	"fmt"
	"net/http"
	"os"

	Client "github.com/dykyi-roman/dive-into-go/client"
	"github.com/dykyi-roman/dive-into-go/server/Domain"
	"github.com/dykyi-roman/dive-into-go/server/Infrastructure"
)

func setupRoutes() {
	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		form := Client.UploadForm{}
		form.FileUploadForm(w, r)
	})

	http.HandleFunc("/api/v1/upload", func(w http.ResponseWriter, r *http.Request) {
		var storage Domain.FileUploadInterface = Infrastructure.FileUploadLocalClient{}
		if storage.FilesUpload(w, r) {
			fmt.Fprintf(w, "Successfully Uploaded File(s)\n")
		}
	})

	http.HandleFunc("/api/v1/db/insert", func(w http.ResponseWriter, r *http.Request) {
		var db Domain.DataBaseInterface = Infrastructure.RedisClient{}
		if db.Insert(r) {
			fmt.Fprintf(w, "Success Insert\n")
		}
	})

	http.HandleFunc("/api/v1/db/get", func(w http.ResponseWriter, r *http.Request) {
		var db Domain.DataBaseInterface = Infrastructure.RedisClient{}

		var payload = db.Get()
		fmt.Println(payload)
	})
}

func main() {
	var PORT string
	if PORT = os.Getenv("GO_PORT"); PORT == "" {
		fmt.Println("Port is not open!")
		os.Exit(1)
	}

	setupRoutes()

	http.ListenAndServe(":"+PORT, nil)
}
