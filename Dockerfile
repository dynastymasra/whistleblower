FROM golang:1.14.0-alpine AS builder

WORKDIR /go/src/github.com/dynastymasra/whistleblower

# install dependecies
RUN set -ex \
    && apk add --update bash git curl \
    && git config --global http.https://gopkg.in.followRedirects true \
    && rm -rf /var/cache/apk/*
COPY go.mod go.sum ./
RUN go mod download

## build linux app source code
COPY . ./
RUN GOOS=linux go build -tags=main -o whistleblower

FROM alpine:3.11.3
RUN set -ex && apk add --update bash ca-certificates tzdata \
    && rm -rf /var/cache/apk/*

# app
WORKDIR /app
COPY --from=builder /go/src/github.com/dynastymasra/whistleblower/whistleblower /app/
COPY --from=builder /go/src/github.com/dynastymasra/whistleblower/migration /app/migration
COPY --from=builder /go/src/github.com/dynastymasra/whistleblower/*.sh /app/

## runtime configs
EXPOSE 8080
ENTRYPOINT ["./whistleblower"]