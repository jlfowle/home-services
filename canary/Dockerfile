FROM golang:1.18-bullseye AS builder

WORKDIR $GOPATH/src/github.com/jlfowle/home-services/canary/
COPY . .

RUN go mod download

ARG GOOS=linux
ARG GOARCH=amd64

RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o /go/bin/canary cmd/main.go

FROM scratch

COPY --from=builder /go/bin/canary /go/bin/canary

EXPOSE 8081/tcp

ENTRYPOINT ["/go/bin/canary"]
