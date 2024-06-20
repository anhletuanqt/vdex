FROM golang:1.21-alpine as builder

RUN apk --no-cache add ca-certificates
RUN apk add --update --no-cache git build-base
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -buildvcs=false -ldflags="-s -w" -o server .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/server", "/build/conf.json", "/"]
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000

ENTRYPOINT ["/server"]
