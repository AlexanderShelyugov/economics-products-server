FROM golang:1.14.2-alpine AS build
WORKDIR /app
COPY go.mod
RUN go mod download
ADD . ./
RUN go build ./cmd/server.go

FROM alpine
WORKDIR /app
COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/server /app/.env /app/
EXPOSE 8100
ENTRYPOINT /app/server