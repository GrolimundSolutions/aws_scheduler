# Obtain certs for final stage
FROM alpine:3.18.4 as authority
RUN mkdir /user && \
    echo 'scheduler:x:1001:1001:scheduler:/:' > /user/passwd && \
    echo 'schedulergroup:x:1001:' > /user/group
RUN apk --no-cache add ca-certificates tzdata

# Build app binary for final stage
FROM golang:1.21.1-alpine3.18 AS builder
RUN apk add git make
WORKDIR /app
COPY . .
ADD https://truststore.pki.rds.amazonaws.com/eu-central-1/eu-central-1-bundle.pem /aws-cert/rds-tls-ca.pem
RUN chown -R 1001:0 /aws-cert/ && \
    chmod -R g+rwx /aws-cert/
RUN go mod download && \
    make build

# Final stage
#FROM alpine:3.15.0
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=authority /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/aws_scheduler ./
COPY --from=builder /aws-cert/rds-tls-ca.pem /aws-cert/rds-tls-ca.pem
COPY /database/PROD_migrations/* ./database/PROD_migrations/
COPY app.env ./app.env
USER scheduler:schedulergroup
ENTRYPOINT ["./aws_scheduler"]