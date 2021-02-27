## BASE IMAGE
FROM golang:1.16.0-alpine3.13

RUN mkdir /ingilizce-kelime-go

## Copy the current directory into our newly created directory.
ADD . /ingilizce-kelime-go

WORKDIR /ingilizce-kelime-go

RUN go mod download

RUN go build -o main .

CMD ["/ingilizce-kelime-go/main"]