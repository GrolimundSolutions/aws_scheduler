# Obtain certs for final stage
FROM alpine:3.15.0 as authority
RUN mkdir /user && \
    echo 'scheduler:x:1000:1000:scheduler:/:' > /user/passwd && \
    echo 'schedulergroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates tzdata

# Build app binary for final stage
FROM golang:1.18.0-alpine3.15 AS builder
RUN apk add git make
WORKDIR /app
COPY . .
RUN go mod download
RUN make build


# Final stage
#FROM alpine:3.15.0
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=authority /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/aws_scheduler ./
COPY /database/PROD_migrations/* ./database/PROD_migrations/
COPY app.env ./app.env
USER scheduler:schedulergroup
ENTRYPOINT ["./aws_scheduler"]