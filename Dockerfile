# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################

FROM golang:1.21.5-alpine3.18 as builder

WORKDIR /app
RUN apk add build-base

COPY . .

RUN go build -buildvcs=false -tags osusergo,netgo -o /usr/local/bin/api ./cmd/api

############################
# STEP 2 build a small image
############################
FROM alpine

COPY --from=builder /usr/local/bin/api /usr/local/bin/api

# command to be used to execute when the image is used to start a container
CMD [ "api" ]
