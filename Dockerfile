FROM golang:1.16.0 AS Builder

WORKDIR /ingilizce-kelime-go

COPY go.sum go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o main .

FROM scratch
WORKDIR /ingilizce-kelime-go
COPY --from=0 /ingilizce-kelime-go/ .
CMD ["/ingilizce-kelime-go/main"]


