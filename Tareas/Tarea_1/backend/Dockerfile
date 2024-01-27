
FROM golang:1.22rc1-bullseye AS build

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o app

FROM scratch

WORKDIR /

COPY --from=build /app/app app

CMD ["/app"]