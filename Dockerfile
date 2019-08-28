FROM golang:1.12 as builder
WORKDIR /module
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux TIME=20190701 go build -o app cmd/main/main.go

FROM alpine:3
WORKDIR /root
COPY --from=builder /module/app /root
COPY /config /root/config
EXPOSE 3000
CMD ["./app"]