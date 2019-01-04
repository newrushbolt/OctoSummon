FROM golang:alpine3.8 as build
COPY . /go/src/github.com/newrushbolt/OctoSummon
WORKDIR /go/src/github.com/newrushbolt/OctoSummon
RUN  apk add --no-cache git make gcc libc-dev ca-certificates \
  && make deps \
  && make

FROM library/alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/newrushbolt/OctoSummon/bin /OctoSummon
CMD ["/OctoSummon"]
