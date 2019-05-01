# Dive into Golang

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

# Go

....