FROM golang:1.18-alpine 

RUN apk add curl

WORKDIR /go-chaos

# Copying module files for building image
COPY go.mod .
COPY go.sum .

# Download modules to local cache
RUN go mod download

# Order of these 2 steps is important
# If go mod tidy is ran before, it removes all
# dependend modules as no source files are present at this point.
COPY . .
RUN go mod tidy

EXPOSE 9010

CMD [ "go", "run", "cmd/main/main.go" ]

