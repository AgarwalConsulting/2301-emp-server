FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED 0
RUN go build -tags netgo -o server ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/server ./
ENV PORT 8000
CMD [ "/app/server" ]
