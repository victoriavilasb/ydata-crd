FROM golang:1.13-alpine3.10 as builder

ENV CGO_ENABLED=0

WORKDIR /ydata-crd

COPY . .
RUN go mod download
RUN go build -o /server -mod=readonly ./cmd/server.go

FROM alpine

COPY --from=builder /server /server

ENTRYPOINT [ "/server" ]