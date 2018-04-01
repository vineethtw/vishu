FROM golang:1.8.3 as builder
WORKDIR /go/src/github.com/vineethtw/vishu
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vishu .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/vineethtw/vishu/vishu .
CMD ["./vishu"]