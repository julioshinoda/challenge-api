FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/api cmd/api/main.go

FROM scratch
COPY --from=build /bin/api /bin/api
ENTRYPOINT ["/bin/api"]