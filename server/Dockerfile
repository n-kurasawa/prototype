# syntax=docker/dockerfile:1
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server
COPY --from=build /app/config.yaml /etc/config/config.yaml

EXPOSE 8080

USER nonroot:nonroot

CMD ["/server", "-config", "/etc/config/config.yaml", "-port", "8080"]
