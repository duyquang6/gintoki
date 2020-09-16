FROM golang:1.15.2-alpine3.12 AS build-env
RUN apk --no-cache add build-base git mercurial gcc
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build-env /app/gintoki /app/
ENTRYPOINT ./gintoki