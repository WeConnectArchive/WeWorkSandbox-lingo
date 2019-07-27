FROM golang:1.12-alpine

RUN apk add --no-cache git

ENV CGO_ENABLED 0

# Set the working directory to /app
WORKDIR /

COPY ./cmd /cmd
COPY ./internal /internal
COPY ./pkg /pkg
COPY build.go .
COPY go.mod .
COPY go.sum .

RUN go get
#RUN mage -v

EXPOSE 80

# Define environment variable
#ENV NAME World

CMD ["ls"]

