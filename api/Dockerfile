FROM golang:1.24

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x /usr/src/app/env.sh

# RUN go build -v -o /usr/local/bin/main .

# CMD ["/bin/bash", "-c", "source /usr/src/app/env.sh && /usr/local/bin/main"]
