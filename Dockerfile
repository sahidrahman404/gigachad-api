# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################

FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o=/usr/local/bin/api ./cmd/api

############################
# STEP 2 build a small image
############################
FROM alpine:3.19

COPY --from=builder /usr/local/bin/api /usr/local/bin/api

# command to be used to execute when the image is used to start a container
CMD [ "api" ]
