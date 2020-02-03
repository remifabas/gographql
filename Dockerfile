

FROM golang:latest AS build-env
ADD . /go/src/github.com/gographql
WORKDIR /go/src/github.com/gographql
# go get and build <-- THIS IS THE IMPORTANT PART
RUN go get -d -v
RUN go build -o /go/bin/main

CMD ["/go/bin/main"]

#FROM scratch 
# <- Second step to build minimal image
#COPY --from=build-env /go/bin/main /go/bin/main
#ENTRYPOINT ["/go/bin/main"]