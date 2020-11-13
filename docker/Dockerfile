##### build
FROM golang:1.14-alpine AS build

RUN apk --no-cache add ca-certificates

# add user
RUN addgroup -g 2000 crawlr
RUN adduser -D -u 1000 -G crawlr crawlr -s /sbin/nologin

RUN mkdir /app
WORKDIR /app

ADD go.mod go.sum ./
ADD cmd ./cmd
ADD internal ./internal
ADD pkg ./pkg
ADD vendor ./vendor

##### compile
FROM build AS compile

RUN CGO_ENABLED=0 GOOS=linux \
    go build -mod=vendor -a -installsuffix cgo \
    -o bin/crawlr \
    ./cmd/main.go

##### start application
FROM scratch

# copy ca certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=compile \
    /etc/passwd \
    /etc/passwd
COPY --from=compile \
    /app/bin/ \
    /usr/local/bin/
USER crawlr

ENTRYPOINT ["/usr/local/bin/crawlr"]