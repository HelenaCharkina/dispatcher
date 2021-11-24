FROM golang:1.16.8

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o dispatcher
CMD ["./dispatcher"]