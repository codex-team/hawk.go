FROM golang:1.15.2
ENV GO111MODULE=on
WORKDIR /test
COPY . .
RUN go build .
EXPOSE 9090
CMD ["./server"]
