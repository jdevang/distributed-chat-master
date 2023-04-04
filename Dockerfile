FROM golang as builder


RUN apt update && apt upgrade && \
    apt install bash git make

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1

RUN go mod tidy
RUN go build -tags netgo -a -v -installsuffix cgo -o bin/master main.go 


FROM alpine:3
RUN apk update \
    && apk add --no-cache curl wget \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /app/bin/master /master

CMD ["/master"]