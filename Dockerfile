# Two-stage build:
#    first  FROM prepares a binary file in full environment ~780MB
#    second FROM takes only binary file ~10MB

FROM golang:1.9 AS builder

RUN go version

COPY . "/go/src/github.com/gographql"
WORKDIR "/go/src/github.com/gographql"

RUN go get -v -t  .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /main

CMD ["/main"]

EXPOSE 8383



#########
# second stage to obtain a very small image
FROM scratch

COPY --from=builder /main .

EXPOSE 8383

CMD ["/main"]