FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY config.json ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /notifynyc
CMD ["/notifynyc"]

