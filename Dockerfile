FROM golang:alpine as builder


RUN apk --no-cache add git
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 go build -mod vendor -o ./apiServer main.go
EXPOSE 8080

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates

WORKDIR /server/
COPY --from=0 /go/src/apiServer .
COPY --from=builder /go/src/runtime /server/runtime
COPY --from=builder /go/src/app/template /server/app/template

ENTRYPOINT ["./apiServer"]

