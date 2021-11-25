# Obtain certs for final stage
FROM alpine:3.15.0 as authority
RUN mkdir /user && \
    echo 'scheduler:x:1000:1000:scheduler:/:' > /user/passwd && \
    echo 'schedulergroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.17.3-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# Final stage
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
USER scheduler:schedulergroup
ENTRYPOINT ["./main"]