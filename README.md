# Dive into Golang

# Documentation 

* [How to Write Go Code](https://golang.org/doc/code.html#GOPATH)

* [Go Tour](https://tour.golang.org/welcome/1)

* [Go Short Guide](https://metanit.com/go/web/)

* [Go Full Guide](https://metanit.com/go/tutorial/)

* [Go Doc](https://golang.org/doc/)

# Go in Visual Studio Code

Using the Go extension for Visual Studio Code, you get language features like IntelliSense, code navigation, symbol search, bracket matching, snippets and many more that will help you in Golang development.

* [Go Plugin](https://github.com/Microsoft/vscode-go)

* https://metanit.com/go/tutorial/1.4.php

* https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a

* https://code.visualstudio.com/docs/languages/go

# Go application with Docker

## Run from Dockerfile

Build image from our Docker file by executing this line at your terminal:

`docker build . -t golang-image`

After running this, a new docker image has been built with the name we provided (golang-image). Now we can run a container from it:

`docker run -p 8080:3001 -it golang-image`

Go to your browser and navigate to localhost:8080

## Run from Docker-compose

`docker-compose up`

Go to your browser and navigate to localhost:8080

# Go Debugging

....

## Author
[Dykyi Roman](https://www.linkedin.com/in/roman-dykyi-43428543/), e-mail: [mr.dukuy@gmail.com](mailto:mr.dukuy@gmail.com)


