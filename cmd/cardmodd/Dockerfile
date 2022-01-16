# build: alpine layer
FROM golang:alpine as build

# build: install build dependencies
RUN apk update \
    && apk add curl git bash make ca-certificates \
    && rm -rf /var/cache/apk/*

# build: install database migration tool
ADD https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin

# build: install modules files
WORKDIR /app
COPY go.* ./
RUN go mod download \
    && go mod verify

# build: install source files & compile
COPY . .
RUN make build-binary

# api: final api layer
FROM alpine:latest

# api: install runtime dependencies
RUN apk update \
    && apk --no-cache add ca-certificates bash \
    && rm -rf /var/cache/apk/*

# api: copy over binaries and entrypoint
WORKDIR /app
COPY --from=build /usr/local/bin/migrate /usr/local/bin
COPY --from=build /app/migrations ./migrations/
COPY --from=build /app/cmd/cardmodd/entrypoint.sh .
COPY --from=build /app/build/gardner .
COPY --from=build /app/build/cardmodd .

# api: create non-root user
RUN addgroup -S cardmod --gid 671 \
    && adduser -S cardmod --uid 671 -G cardmod \
    && chown -R cardmod:cardmod /app

# api: start the service
USER cardmod
ENTRYPOINT [ "./entrypoint.sh" ]