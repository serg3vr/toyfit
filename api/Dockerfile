FROM golang:1.24

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x /usr/src/app/env.sh
# RUN go build -v -o /usr/local/bin/app ./...
RUN go build -v -o /usr/local/bin/main .

# CMD ["app"]
# CMD ["/usr/local/bin/app"]
CMD ["/bin/bash", "-c", "source /usr/src/app/env.sh && /usr/local/bin/main"]
