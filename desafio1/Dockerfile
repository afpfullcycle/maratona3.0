FROM golang:rc-alpine3.12 AS buildimage
WORKDIR /app/build
COPY server* .
RUN go build main.go
RUN ls -la

FROM alpine:3.12.0 AS app
WORKDIR /app
COPY --from=buildimage /app/build/main .
COPY template* .
RUN ls -la
ENTRYPOINT [ "./main" ]
