FROM golang:1.22.12-alpine AS builder 

COPY . /go/src/github.com/irootpro/chat-auth
WORKDIR /go/src/github.com/irootpro/chat-auth

RUN go mod download
RUN go build -o chat-auth ./cmd/main.go

FROM alpine:3.21
COPY --from=builder /go/src/github.com/irootpro/chat-auth/chat-auth .
CMD ["./chat-auth"]
 

