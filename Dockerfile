FROM golang:1.16.0 AS Builder
WORKDIR /ingilizce-kelime-go
COPY go.sum go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.7
RUN apk add --no-cache curl
#HEALTHCHECK --interval=2s --timeout=10s \
 #   CMD curl -f 127.0.0.1:8080/healtcheck || exit 1
WORKDIR /ingilizce-kelime-go
COPY --from=Builder /ingilizce-kelime-go/ .
CMD ["/ingilizce-kelime-go/main"]


