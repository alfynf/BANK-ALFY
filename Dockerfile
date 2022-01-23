FROM golang:1.17-alpine
FROM mysql:latest
WORKDIR /bank-alfy
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o program
CMD ./program