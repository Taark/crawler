FROM golang:1.15.6 as builder
COPY ./src /go/src/web-crawler
WORKDIR /go/src/web-crawler
ENV GOPATH /go
ENV GO111MODULE "on"
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web-crawler

FROM scratch as web-crawler
COPY --from=builder /go/src/web-crawler/web-crawler /usr/local/bin/web-crawler
ENTRYPOINT ["/usr/local/bin/web-crawler"]