FROM golang:1.22.2-alpine3.19 as build
WORKDIR /tmp
COPY . /tmp
RUN go build -o /var/app main.go

FROM alpine:3.19
LABEL org.opencontainers.image.source="https://github.com/glieske/kube-nanny"
RUN mkdir -pv /app
WORKDIR /app
COPY --from=build /var/app /app/app
CMD ["/app/app"]
