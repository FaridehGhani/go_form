FROM golang:1.18

WORKDIR /go/src/github.com/FaridehGhani/go_form

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /go/src/github.com/FaridehGhani/go_form ./...

CMD ["go_form"]