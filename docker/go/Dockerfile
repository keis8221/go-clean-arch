# ===== build go binary =====
FROM golang:1.22.0 AS go-builder

WORKDIR /go/src/app/

COPY cmd/rest/main.go .
COPY go.mod .
COPY go.sum .
RUN mkdir pkg
COPY pkg/ ./pkg

RUN go mod download

RUN go build -o server main.go

# ===== build docker image =====
FROM alpine

RUN apk --no-cache add tzdata

COPY --from=go-builder /go/src/app/server server

ENTRYPOINT ["/server"]
