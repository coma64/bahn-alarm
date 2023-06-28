FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app
ENV BIN_DIR=/app/bin

RUN apk add upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p $BIN_DIR \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $BIN_DIR/server \
    && upx --best --lzma $BIN_DIR/server

# build the cmds
RUN cd $BIN_DIR \
    && find /app/cmds/ -type f \
    | xargs -n1 -I {} sh -c 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $BIN_DIR/$(basename {} .go) {} && upx --best --lzma $BIN_DIR/$(basename {} .go)'

# FROM gcr.io/distroless/base-debian11
# a shell is pretty helpful :)
FROM alpine:3.18.0

WORKDIR /app
ENV BIN_DIR=/app/bin
ENV PATH="/app/bin:$PATH"

# need this for timezone conversions
RUN apk add tzdata

COPY --from=builder $BIN_DIR $BIN_DIR
COPY --from=builder /app/migrations /app/migrations
COPY config.yml config.dev.yml ./
COPY scripts/ ./scripts/

EXPOSE 8090
USER nobody:nobody

CMD $BIN_DIR/server