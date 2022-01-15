# build: alpine layer
FROM golang:alpine as build

# build: install build dependencies
RUN apk update \
    && apk add curl git bash make ca-certificates \
    && rm -rf /var/cache/apk/*

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

# api: copy over binary
WORKDIR /app
COPY --from=build /app/cmd/cardmodd/entrypoint.sh .
COPY --from=build /app/build/cardmodd .

# api: create non-root user
RUN addgroup -S cardmod --gid 671 \
    && adduser -S cardmod --uid 671 -G cardmod \
    && chown -R cardmod:cardmod /app

# api: start the service
USER cardmod
ENTRYPOINT [ "./entrypoint.sh" ]