FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app
ENV BIN=/app/bahn-alarm

RUN apk add upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $BIN \
    && upx --best --lzma $BIN

# FROM gcr.io/distroless/base-debian11
# a shell is pretty helpful :)
FROM alpine:3.18.0

WORKDIR /app
ENV BIN=/app/bahn-alarm
ENV PATH="/app:$PATH"

# need this for timezone conversions
RUN apk add tzdata

COPY --from=builder $BIN $BIN
COPY --from=builder /app/migrations /app/migrations
COPY config.yml config.dev.yml ./
COPY scripts/ ./scripts/

EXPOSE 8090
USER nobody:nobody

CMD $BIN serve