FROM golang:1.20-bullseye AS builder

WORKDIR $GOPATH/src/github.com/jlfowle/home-services/zilch/
COPY . .

RUN go mod download

ARG GOOS=linux
ARG GOARCH=amd64

RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o /go/bin/zilch cmd/main.go

FROM scratch

COPY --from=builder /go/bin/zilch /go/bin/zilch

USER 1001

EXPOSE 8080/tcp

ENTRYPOINT ["/go/bin/canary"]
