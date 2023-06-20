FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir /app/bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/bin/server

# build the commands
RUN cd /app/bin && find ../cmds/ -type f | xargs -n1 -I {} sh -c 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" {}'

# FROM gcr.io/distroless/base-debian11
# a shell is pretty helpful :)
FROM alpine:3.18.0

WORKDIR /app
ENV PATH="/app/bin:$PATH"

COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/migrations /app/migrations
COPY config.yml config.dev.yml ./

EXPOSE 8080
USER nobody:nobody

CMD ["/app/bin/server"]