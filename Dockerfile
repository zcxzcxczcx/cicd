FROM golang:alpine
WORKDIR /second-website 
COPY . /second-website 
entrypoint go run main.go