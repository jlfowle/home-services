FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/jlfowle/home-services/canary/
COPY . .

RUN go get -d -v

ARG GOOS=linux
ARG GOARCH=amd64

RUN GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w -s" -o /go/bin/canary

FROM scratch

COPY --from=builder /go/bin/canary /go/bin/canary

ENTRYPOINT ["/go/bin/canary"]
