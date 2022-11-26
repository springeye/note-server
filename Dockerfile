FROM golang:1.19-alpine
WORKDIR /app
COPY . .
ENV GOPROXY https://goproxy.cn
RUN go mod tidy
RUN go build -o note-server
CMD ["/app/note-server"]