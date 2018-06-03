FROM golang:alpine as BASE
WORKDIR /go/src/github.com/cmh2166/elag18apis/taquito
COPY . .
RUN apk update && apk add --no-cache ca-certificates && \
    apk add --no-cache --virtual .build-deps git && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    apk del .build-deps
RUN CGO_ENABLED=0 GOOS=linux go build -o api -ldflags "-s" -a -installsuffix cgo main.go

FROM scratch
EXPOSE 8080
COPY --from=BASE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BASE /go/src/github.com/cmh2166/elag18apis/taquito/api .
CMD ["./api"]
