# Build environment
# -----------------
FROM golang:1.15-alpine as build-env
WORKDIR /versla

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate


# Deployment environment
# ----------------------
FROM alpine
RUN apk update

COPY --from=build-env /versla/bin/app /versla/
COPY --from=build-env /versla/bin/migrate /versla/
COPY --from=build-env /versla/migrations /versla/migrations

EXPOSE 8080
CMD ["/versla/app"]
