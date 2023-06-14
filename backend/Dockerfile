FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bahn-alarm

FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Static files
COPY openapi.yml ./
COPY static/ ./static
COPY --from=builder /app/swagger-ui/dist ./swagger-ui/dist

COPY --from=builder /app/bahn-alarm .

EXPOSE 8090
USER nonroot:nonroot

ENTRYPOINT ["/app/bahn-alarm"]