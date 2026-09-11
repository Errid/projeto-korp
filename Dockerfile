FROM golang:1.25 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY app/ ./app/

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -o /http-server-projeto-korp \
    ./app
FROM scratch

COPY --from=build /http-server-projeto-korp /http-server-projeto-korp

USER 10001:10001

EXPOSE 8080

ENTRYPOINT ["/http-server-projeto-korp"]