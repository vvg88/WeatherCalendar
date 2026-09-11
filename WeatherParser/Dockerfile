FROM golang:1.27.1
WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go install -v ./...

EXPOSE 8080
ENV TZ Europe/Moscow

CMD ["WeatherParser"]
