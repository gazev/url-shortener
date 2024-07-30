FROM golang:1.22.5 AS build-stage

WORKDIR /mus

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux CGO_ENABLED=0 go build -o exe

FROM alpine:3.14 AS build-release-stage

WORKDIR /

COPY --from=build-stage /mus/exe /mus

EXPOSE 8000

ENTRYPOINT ["/mus"]
