FROM golang:alpine
ENV GOPROXY=https://goproxy.cn
WORKDIR /second-website 
COPY . /second-website 
entrypoint go run main.go