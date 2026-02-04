FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /out/items-api ./main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /out/items-api /app/items-api
COPY etc /app/etc
USER nonroot:nonroot
EXPOSE 8888
ENTRYPOINT ["/app/items-api", "-f", "/app/etc/items-api.yaml"]
