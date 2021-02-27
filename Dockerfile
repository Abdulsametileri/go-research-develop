## BASE IMAGE
FROM golang:1.16.0-alpine3.13

## v1: RUN mkdir /ingilizce-kelime-go

## Copy the current directory into our newly created directory.
## v1: ADD . /ingilizce-kelime-go

WORKDIR /ingilizce-kelime-go
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

CMD ["/ingilizce-kelime-go/main"]