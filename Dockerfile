FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /server /server
ENV ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/server"]
